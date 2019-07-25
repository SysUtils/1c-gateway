package resolver

import (
	"github.com/essentialkaos/translit"
	"strings"
	"unicode"
)

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
		src = "[]" + g.translateType(src[11:len(src)-1])
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
	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}
