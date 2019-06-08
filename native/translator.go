package native

import (
	"github.com/essentialkaos/translit"
	"strings"
)

func (g *Generator) TranslateType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		src = "[]" + g.TranslateType(src[11:len(src)-1])
	}
	if val, ok := g.TypeMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}

func (g *Generator) TranslateName(src string) string {
	if val, ok := g.NameMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}
