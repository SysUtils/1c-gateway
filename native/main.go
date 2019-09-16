// Package native provides generator for native odata client
package native

import (
	"fmt"
	generated "github.com/SysUtils/1c-gateway"
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

func extractAsset(asset, path string) {
	data, err := generated.Asset(asset)
	if err != nil {
		log.Panic(err)
	}
	f, err := os.Create(path)
	if err != nil {
		log.Panic(err)
	}
	_, err = f.Write(data)
	if err != nil {
		log.Panic(err)
	}
	err = f.Close()
	if err != nil {
		log.Panic(err)
	}
}

// NewGenerator returns initialized generator
func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string), Associations: make(map[string]map[string]string)}
}

// Generate generates the native odata client and writes it to ./odata folder
func (g *Generator) Generate() {
	// static/binary.go
	// static/boolean.go
	// static/client.go
	// static/client_entity.go
	// static/datetime.go
	// static/double.go
	// static/errors.go
	// static/float.go
	// static/guid.go
	// static/int.go
	// static/int16.go
	// static/int64.go
	// static/interfaces.go
	// static/stream.go
	// static/string.go
	// static/time.go
	// static/where.go
	extractAsset("static/binary.go", "odata/binary.go")
	extractAsset("static/boolean.go", "odata/boolean.go")
	extractAsset("static/client.go", "odata/client.go")
	extractAsset("static/client_entity.go", "odata/client_entity.go")
	extractAsset("static/datetime.go", "odata/datetime.go")
	extractAsset("static/double.go", "odata/double.go")
	extractAsset("static/errors.go", "odata/errors.go")
	extractAsset("static/float.go", "odata/float.go")
	extractAsset("static/guid.go", "odata/guid.go")
	extractAsset("static/int.go", "odata/int.go")
	extractAsset("static/int16.go", "odata/int16.go")
	extractAsset("static/int64.go", "odata/int64.go")
	extractAsset("static/interfaces.go", "odata/interfaces.go")
	extractAsset("static/stream.go", "odata/stream.go")
	extractAsset("static/string.go", "odata/string.go")
	extractAsset("static/time.go", "odata/time.go")
	extractAsset("static/where.go", "odata/where.go")
	extractAsset("static/grpc_helper.go", "odata/grpc_helper.go")

	data := fmt.Sprintf(`package odata
import (
"encoding/json"
"errors"
)

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
