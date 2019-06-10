// Package native provides generator for native odata client
package native

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
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string), Associations: make(map[string]map[string]string)}
}

// Generate generates the native odata client and writes it to ./odata folder
func (g *Generator) Generate() {
	data := fmt.Sprintf(`package odata
import "encoding/json"

%s`, g.genTypes(g.schema.Entities))
	g.writeGofile("Entity.go", data)

	data = fmt.Sprintf(`package odata

%s`, g.genPrimaryKeys(g.schema.Entities))
	g.writeGofile("Primary.go", data)

	data = fmt.Sprintf(`package odata

%s`, g.genComplexTypes(g.schema.Complexes))
	g.writeGofile("Complex.go", data)

	data = fmt.Sprintf(`package odata

%s`, g.genFunctions(g.schema.Functions))
	g.writeGofile("Functions.go", data)

	data = fmt.Sprintf(`package odata

%s`, g.genNavigations(g.schema.Entities))
	g.writeGofile("Navigations.go", data)

	data = fmt.Sprintf(`package odata

%s`, g.genFilters(g.schema.Entities))
	g.writeGofile("Filters.go", data)
}

func (g *Generator) writeGofile(filename, data string) {
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
