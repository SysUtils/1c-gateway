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
	f, _ := os.Create("odata/GrpcConverters.go")
	f.WriteString(`package odata
`)
	f.WriteString(g.GenConverters(g.schema.Entities))
	f.Close()

}
