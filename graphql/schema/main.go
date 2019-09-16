// Package graphql/schema provides generator for GraphQL gateway's schema
package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"log"
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

// Generate generates the GraphQL schema and writes it to ./odata folder
func (g *Generator) Generate() {
	g.extractAssociations(g.schema.Association)
	data := fmt.Sprintf(`package odata

const gqlSchema = `+"`"+`schema {
  query: Query
  mutation: Mutation
  subscription: Subscription
}

input Options {
	Top: Int!
	Skip: Int!
}

scalar String
scalar Int
scalar Int16
scalar Int64
scalar Double
scalar Float
scalar DateTime
scalar Binary
scalar Stream
scalar Boolean
scalar Guid

%s
%s
%s
%s
%s
%s
%s
%s
%s
%s`+"`",
		g.genTypes(g.schema.Entities),
		g.genTypes(g.schema.Complexes),
		g.genPrimaryKeys(g.schema.Entities),
		g.genMutations(g.schema.Entities, g.schema.Functions),
		g.genQueries(g.schema.Entities),
		g.genSubscriptions(g.schema.Entities),
		g.genFilters(g.schema.Entities),
		g.genFilters(g.schema.Complexes),
		g.genFields(g.schema.Entities),
		g.genFunctions(g.schema.Functions))
	g.writeFile("Schema.go", data)
}
func (g *Generator) writeFile(filename, data string) {
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
