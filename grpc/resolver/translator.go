package resolver

import (
	"github.com/essentialkaos/translit"
	"strings"
	"unicode"
)

var ScalarTypes = map[string]string{
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

func ToCamelCase(src string) string {
	result := []rune(src)
	result[0] = unicode.ToUpper(result[0])
	return string(result)
}

func ToLowerCamelCase(src string) string {
	result := []rune(src)
	result[0] = unicode.ToLower(result[0])
	return string(result)
}

func (g *Generator) translateGrpcType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		return "[]" + g.translateGrpcType(src[11:len(src)-1])
	}
	if val, ok := g.TypeMap[src]; ok {
		return ToCamelCase(val + "Grpc")
	}
	if val, ok := ScalarTypes[src]; ok {
		return ToCamelCase(val)
	}
	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)) + "Grpc")
}

func (g *Generator) translateNativeType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		return "[]" + g.translateNativeType(src[11:len(src)-1])
	}
	if val, ok := g.TypeMap[src]; ok {
		return ToCamelCase(val)
	}
	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}

func (g *Generator) translateName(src string) string {
	if val, ok := g.NameMap[src]; ok {
		return ToCamelCase(val)
	}
	if strings.HasSuffix(src, "_key") {
		return g.translateName(src[:len(src)-4]) + "Key"
	}

	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}
