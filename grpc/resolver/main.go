package resolver

import (
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
	"os"
)

type Generator struct {
	TypeMap      map[string]string
	NameMap      map[string]string
	Associations map[string]map[string]string
	schema       shared.Schema
}

func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string), Associations: map[string]map[string]string{}}
}

func (g *Generator) Start() {
	f, _ := os.Create("odata/GrpcConverter.go")
	f.WriteString(`package odata
`)
	f.WriteString(`func (w *BaseWhere) ToNative() Where {
	return Where {Top: w.Top,
	Skip: w.Skip,
	Orderby: w.Orderby }
}
`)

	f.WriteString(g.GenTypeConverters(g.schema.Entities))
	f.WriteString(g.GenComplexConverters(g.schema.Complexes))
	f.Close()

	f, _ = os.Create("odata/GrpcResolver.go")
	f.WriteString(`package odata

import "context"

type GrpcResolver struct {
	Client *Client
}

`)
	f.WriteString(g.GenQueryResolvers(g.schema.Entities))
	f.Close()
}
