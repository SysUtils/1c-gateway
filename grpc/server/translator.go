package server

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

func (g *Generator) translateName(src string) string {
	if val, ok := g.NameMap[src]; ok {
		return ToCamelCase(val)
	}
	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}
