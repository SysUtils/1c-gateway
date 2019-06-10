package server

import (
	"log"
	"regexp"
)

func (g *Generator) extractNameMap() map[string]string {
	result := make(map[string]string)
	for _, e := range g.schema.Entities {
		result[g.translateName(e.Name)] = e.Name
		for _, p := range e.Properties {
			result[g.translateName(p.Name)] = p.Name
		}
	}

	for _, e := range g.schema.Complexes {
		result[g.translateName(e.Name)] = e.Name
		for _, p := range e.Properties {
			result[g.translateName(p.Name)] = p.Name
		}
	}
	return result
}

func (g *Generator) replacer(s string) string {
	s = s[6:]
	return `json:"` + g.ReverseNameMap[s]
}

func (g *Generator) fixJson(source string) string {
	g.ReverseNameMap = g.extractNameMap()
	reg, err := regexp.Compile(`json:"[A-z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	res := reg.ReplaceAllStringFunc(source, g.replacer)
	return res
}
