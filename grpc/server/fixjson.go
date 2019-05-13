package server

import (
	"log"
	"regexp"
)

func (g *Generator) ExtractNameMap() map[string]string {
	result := make(map[string]string)
	for _, e := range g.schema.Entities {
		result[g.TranslateName(e.Name)] = e.Name
		for _, p := range e.Properties {
			result[g.TranslateName(p.Name)] = p.Name
		}
	}

	for _, e := range g.schema.Complexes {
		result[g.TranslateName(e.Name)] = e.Name
		for _, p := range e.Properties {
			result[g.TranslateName(p.Name)] = p.Name
		}
	}
	return result
}

func (g *Generator) Replacer(s string) string {
	s = s[6:]
	return `json:"` + g.ReverseNameMap[s]
}

func (g *Generator) FixJson(source string) string {
	g.ReverseNameMap = g.ExtractNameMap()
	reg, err := regexp.Compile(`json:"[A-z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	res := reg.ReplaceAllStringFunc(source, g.Replacer)
	return res
}
