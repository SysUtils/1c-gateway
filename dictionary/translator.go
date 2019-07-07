package dictionary

import (
	"github.com/essentialkaos/translit"
	"strings"
)

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
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}

func (g *Generator) translateName(src string) string {
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}
