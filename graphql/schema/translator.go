package schema

import (
	"pkg.re/essentialkaos/translit.v2"
	"strings"
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

func (g *Generator) TranslateType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		src = "[" + g.TranslateType(src[11:len(src)-1]) + "!]"
	}
	if val, ok := g.TypeMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}

func (g *Generator) TranslateInputType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		return "[" + g.TranslateInputType(src[11:len(src)-1]) + "!]"
	}
	if val, ok := g.TypeMap[src]; ok {
		src = val
	}
	if _, ok := ScalarTypes[src]; !ok {
		src += "Input"
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}

func (g *Generator) TranslateName(src string) string {
	if val, ok := g.NameMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}
