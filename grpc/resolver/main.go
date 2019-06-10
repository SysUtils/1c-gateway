// Package grpc/resolver provides generator for GRPC resolver
package resolver

import (
	"github.com/SysUtils/1c-gateway/shared"
	"os"
)

type Generator struct {
	TypeMap      map[string]string
	NameMap      map[string]string
	Associations map[string]map[string]string
	schema       shared.Schema
}

// NewGenerator returns initialized generator
func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string), Associations: map[string]map[string]string{}}
}

// Generate generates the resolver and writes it to ./odata folder
func (g *Generator) Generate() {
	f, _ := os.Create("odata/GrpcConverter.go")
	f.WriteString(`package odata
`)
	f.WriteString(`func (w *BaseWhere) ToNative() Where {
	return Where {Top: w.Top,
	Skip: w.Skip,
	Orderby: w.Orderby }
}
`)

	f.WriteString(g.genTypeConverters(g.schema.Entities))
	f.WriteString(g.genComplexConverters(g.schema.Complexes))
	f.Close()

	f, _ = os.Create("odata/GrpcResolver.go")
	f.WriteString(`package odata

import "context"

type GrpcResolver struct {
	Client *Client
}

`)
	f.WriteString(g.genQueryResolvers(g.schema.Entities))
	f.Close()
}
