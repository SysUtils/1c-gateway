// Package grpc provides generator for GRPC gateway
package grpc

import (
	"github.com/SysUtils/1c-gateway/grpc/resolver"
	"github.com/SysUtils/1c-gateway/grpc/schema"
	"github.com/SysUtils/1c-gateway/grpc/server"
	"github.com/SysUtils/1c-gateway/shared"
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

// Generate generates the grpc gateway and writes it to ./odata folder
func (g *Generator) Generate() {
	schemaGen := schema.NewGenerator(g.schema)
	schemaGen.TypeMap = g.TypeMap
	schemaGen.NameMap = g.NameMap
	schemaGen.Generate()

	resolverGen := resolver.NewGenerator(g.schema)
	resolverGen.TypeMap = g.TypeMap
	resolverGen.NameMap = g.NameMap
	resolverGen.Generate()

	serverGen := server.NewGenerator(g.schema)
	serverGen.TypeMap = g.TypeMap
	serverGen.NameMap = g.NameMap
	serverGen.Start()
}
