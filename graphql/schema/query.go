package schema

import "gitlab.com/zullpro/core/1cclientgenerator.git/shared"

func (g *Generator) GenQueries(source []shared.OneCType) string {
	result := "type Query {\n"
	for _, entity := range source {
		result += g.GenQuery(entity)
		result += "\n"
	}
	result += "}"

	return result
}

func (g *Generator) GenQuery(source shared.OneCType) string {
	result := ""
	result += "	" + g.TranslateType(source.Name) + "(Key: Primary" + g.TranslateType(source.Name) + "!): " + g.TranslateType(source.Name) + "\n"
	result += "	" + g.TranslateType(source.Name) + "s(): " + g.TranslateType(source.Name)
	return result
}
