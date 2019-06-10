package server

import (
	"github.com/essentialkaos/translit"
	"strings"
)

func (g *Generator) translateName(src string) string {
	if val, ok := g.NameMap[src]; ok {
		return val
	}
	return translit.EncodeToICAO(strings.Replace(src, "_", "", -1))
}
