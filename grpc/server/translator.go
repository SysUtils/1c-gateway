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
	if strings.HasSuffix(src, "_key") {
		return g.translateName(src[:len(src)-4]) + "Key"
	}

	return ToCamelCase(translit.EncodeToICAO(strings.Replace(src, "_", "", -1)))
}
