// Package graphql provides generator for GraphQL gateway's
package graphql

import (
	generated "github.com/SysUtils/1c-gateway"
	"github.com/SysUtils/1c-gateway/graphql/resolver"
	"github.com/SysUtils/1c-gateway/graphql/schema"
	"github.com/SysUtils/1c-gateway/shared"
	"log"
	"os"
)

type Generator struct {
	TypeMap map[string]string
	NameMap map[string]string
	schema  shared.Schema
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

func NewGenerator(schema shared.Schema) *Generator {
	return &Generator{schema: schema, TypeMap: make(map[string]string), NameMap: make(map[string]string)}
}

func (g *Generator) Start() {
	extractAsset("static/server.go", "odata/server.go")
	schemaGen := schema.NewGenerator(g.schema)
	schemaGen.TypeMap = g.TypeMap
	schemaGen.NameMap = g.NameMap
	schemaGen.Generate()

	resolverGen := resolver.NewGenerator(g.schema)
	resolverGen.TypeMap = g.TypeMap
	resolverGen.NameMap = g.NameMap
	resolverGen.Generate()
}
