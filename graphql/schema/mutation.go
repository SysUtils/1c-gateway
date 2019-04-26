package schema

import "gitlab.com/zullpro/core/1cclientgenerator.git/shared"

func (g *Generator) GenMutations(source []shared.OneCType) string {
	result := "type Mutation {\n"
	for _, entity := range source {
		result += g.GenMutation(entity)
		result += "\n"
	}
	result += "}"

	return result
}

func (g *Generator) GenMutation(source shared.OneCType) string {
	result := ""
	result += "	Update" + g.TranslateType(source.Name) + "(Key: Primary" + g.TranslateType(source.Name) + "," + "Entity: " + g.TranslateType(source.Name) + "): " + g.TranslateType(source.Name) + "\n"
	result += "	Remove" + g.TranslateType(source.Name) + "(Key: Primary" + g.TranslateType(source.Name) + "): " + g.TranslateType(source.Name)
	return result
}
