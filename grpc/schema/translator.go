package schema

import (
	"github.com/essentialkaos/translit"
	"strings"
)

var ScalarTypes = map[string]bool{
	"string":    true,
	"int32":     true,
	"int64":     true,
	"double":    true,
	"float":     true,
	"Timestamp": true,
	"bytes":     true,
	"bool":      true,
}

var ScalarTypesMap = map[string]string{
	"String":   "string",
	"Int":      "int32",
	"Int16":    "int32",
	"Int64":    "int64",
	"Double":   "double",
	"Float":    "float",
	"DateTime": "Timestamp",
	"Binary":   "bytes",
	"Stream":   "bytes",
	"Boolean":  "bool",
	"Guid":     "string",
}

func (g *Generator) translateType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		return "repeated " + g.translateType(src[11:len(src)-1])
	}
	if val, ok := g.TypeMap[src]; ok {
		return val + "Grpc"
	}
	if val, ok := ScalarTypesMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1)) + "Grpc"
}

func (g *Generator) translateNativeType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		return "repeated " + g.translateType(src[11:len(src)-1])
	}
	if val, ok := g.TypeMap[src]; ok {
		return val
	}
	if val, ok := ScalarTypesMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}

func (g *Generator) translateName(src string) string {
	if val, ok := g.NameMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}
