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

func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string), Associations: map[string]map[string]string{}}
}

func (g *Generator) writeGqlfile(filename, data string) {
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
	g.ExtractAssociations(g.schema.Association)
	data := fmt.Sprintf(`schema {
  query: Query
  mutation: Mutation
}

input BaseWhere {
	Top: Int!
	Skip: Int!
	Orderby: String!
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
%s`, g.GenTypes(g.schema.Entities),
		g.GenTypes(g.schema.Complexes),
		g.GenPrimaryKeys(g.schema.Entities),
		g.GenMutations(g.schema.Entities),
		g.GenQueries(g.schema.Entities),
		g.GenFilters(g.schema.Entities),
		g.GenFilters(g.schema.Complexes))
	g.writeGqlfile("Schema.graphql", data)
}
