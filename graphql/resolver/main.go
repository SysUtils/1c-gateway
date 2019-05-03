package resolver

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
	"log"
	"os"
)

type Generator struct {
	TypeMap map[string]string
	NameMap map[string]string
	schema  shared.Schema
}

func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string)}
}

func (g *Generator) writeGofile(filename, data string) {
	f, err := os.Create("odata/" + filename)
	if err != nil {
		log.Panic(err)
	}
	_, err = f.WriteString(data)
	if err != nil {
		log.Panic(err)
	}

	err = f.Close()
	if err != nil {
		log.Panic(err)
	}
}

func (g *Generator) Start() {
	data := fmt.Sprintf(`package odata

%s`, g.GenArgs(g.schema.Entities))
	g.writeGofile("Resolver_args.go", data)

	data = fmt.Sprintf(
		`package odata

import "context"

type GqlResolver struct {
	Client *Client
}

%s
%s`, g.GenResolvers(g.schema.Entities), g.GenMutations(g.schema.Entities))
	g.writeGofile("Resolver.go", data)

	data = fmt.Sprintf(`package odata

%s`, g.GenFilters(g.schema.Entities))
	g.writeGofile("Resolver_filter.go", data)
}
