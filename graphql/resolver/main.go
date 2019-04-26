package resolver

import (
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

func (g *Generator) Start() {
	f, err := os.Create("odata/Resolver_args.go")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("package odata\n\n")
	f.WriteString(g.GenArgs(g.schema.Entities))
	f.Close()

	f, err = os.Create("odata/Resolver.go")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(`package odata
import "context"

type GqlResolver struct {
	Client *Client
}

`)
	f.WriteString(g.GenResolvers(g.schema.Entities))
	f.Close()
}
