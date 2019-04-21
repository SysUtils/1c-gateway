package client

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
	f, err := os.Create("odata/Entity.go")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(`package odata
import "encoding/json"
`)
	for _, e := range g.schema.Entities {
		f.WriteString(g.GenType(e) + "\n")
		f.WriteString(g.GenPrimaryKey(e) + "\n")
	}
	f.Close()

	f, err = os.Create("odata/Complex.go")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(`package odata
`)
	for _, e := range g.schema.Complexes {
		f.WriteString(g.GenComplexType(e) + "\n")
	}
	f.Close()
}
