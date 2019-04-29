package schema

import (
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
	"strings"
)

func (g *Generator) GenFilters(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenFilter(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenFilter(source shared.OneCType) string {
	result := ""
	name := g.TranslateType(source.Name) + "Filter"
	result += "input " + name + " {\n"
	result += "	AND: [" + name + "!]\n"
	result += "	OR: [" + name + "!]\n"
	for _, prop := range source.Properties {
		transType := g.TranslateType(prop.Type)
		if _, ok := ScalarTypes[transType]; ok {
			result += "	" + g.TranslateName(prop.Name) + "_eq: " + transType + "\n"
			result += "	" + g.TranslateName(prop.Name) + "_ne: " + transType + "\n"
		} else {
			if !strings.HasPrefix(transType, "[") {
				result += "	" + g.TranslateName(prop.Name) + ": " + transType + "Filter\n"
			}
		}
	}
	result += "}"
	return result
}
