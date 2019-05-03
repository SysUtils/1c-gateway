package grpc

import (
	"gitlab.com/zullpro/core/1cclientgenerator.git/grpc/resolver"
	"gitlab.com/zullpro/core/1cclientgenerator.git/grpc/schema"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
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
	schemaGen := schema.NewGenerator(g.schema)
	schemaGen.TypeMap = g.TypeMap
	schemaGen.NameMap = g.NameMap
	schemaGen.Start()
	resolverGen := resolver.NewGenerator(g.schema)
	resolverGen.TypeMap = g.TypeMap
	resolverGen.NameMap = g.NameMap
	resolverGen.Start()
}
