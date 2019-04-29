package schema

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

	f, err := os.Create("odata/Types.gql")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range g.schema.Entities {
		f.WriteString(g.GenType(e))
	}
	f.Close()

	f, err = os.Create("odata/ComplexTypes.gql")
	for _, e := range g.schema.Complexes {
		f.WriteString(g.GenType(e))
	}
	f.Close()

	f, err = os.Create("odata/PrimaryKeys.gql")
	f.WriteString(g.GenPrimaryKeys(g.schema.Entities))
	f.Close()

	f, err = os.Create("odata/Mutations.gql")
	f.WriteString(g.GenMutations(g.schema.Entities))
	f.Close()

	f, err = os.Create("odata/Queries.gql")
	f.WriteString(g.GenQueries(g.schema.Entities))
	f.Close()

	f, err = os.Create("odata/Filters.gql")
	f.WriteString(g.GenFilters(g.schema.Entities))
	f.WriteString("\n")
	f.WriteString(g.GenFilters(g.schema.Complexes))
	f.Close()

}
