// Package graphql/resolver provides generator for GraphQL gateway's resolver
package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"log"
	"os"
)

type Generator struct {
	TypeMap map[string]string
	NameMap map[string]string
	schema  shared.Schema
}

// NewGenerator returns initialized generator
func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string)}
}

// Generate generates the GraphQL resolver and writes it to ./odata folder
func (g *Generator) Generate() {
	data := fmt.Sprintf(`package odata

%s`, g.genArgs(g.schema.Entities))
	g.writeGofile("Resolver_args.go", data)

	data = fmt.Sprintf(
		`package odata

import "context"

type GqlResolver struct {
	Client *Client
}

%s
%s`, g.genResolvers(g.schema.Entities), g.genMutations(g.schema.Entities))
	g.writeGofile("Resolver.go", data)
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
