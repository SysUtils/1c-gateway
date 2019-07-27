package schema

import (
	"github.com/essentialkaos/translit"
	"strings"
	"unicode"
)

var ScalarTypes = map[string]bool{
	"String":   true,
	"Int":      true,
	"Int16":    true,
	"Int64":    true,
	"Double":   true,
	"Float":    true,
	"DateTime": true,
	"Binary":   true,
	"Stream":   true,
	"Boolean":  true,
	"Guid":     true,
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

func (g *Generator) translateType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		src = "[" + g.translateType(src[11:len(src)-1]) + "!]"
	}
	if val, ok := g.TypeMap[src]; ok {
		return ToCamelCase(val)
	}

	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}

func (g *Generator) translateInputType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		return "[" + g.translateInputType(src[11:len(src)-1]) + "!]"
	}
	if val, ok := g.TypeMap[src]; ok {
		src = val
	}
	if _, ok := ScalarTypes[src]; !ok {
		src += "Input"
	}

	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}

func (g *Generator) translateName(src string) string {
	if val, ok := g.NameMap[src]; ok {
		return ToLowerCamelCase(val)
	}
	if strings.HasSuffix(src, "_key") {
		return g.translateName(src[:len(src)-4]) + "Key"
	}

	return ToLowerCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}
