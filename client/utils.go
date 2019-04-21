package client

import (
	"pkg.re/essentialkaos/translit.v2"
	"strings"
)

func (g *Generator) TranslateType(src string) string {
	src = strings.Replace(src, "_", "", -1)
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		src = "[]" + g.TranslateType(src[11:len(src)-1])
	}
	return translit.EncodeToICAO(src)
}

func (g *Generator) TranslateName(src string) string {
	src = strings.Replace(src, "_", "", -1)
	return translit.EncodeToICAO(src)
}
