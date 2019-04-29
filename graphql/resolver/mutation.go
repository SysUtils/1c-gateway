package resolver

import "gitlab.com/zullpro/core/1cclientgenerator.git/shared"

func (g *Generator) GenMutations(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenMutation(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenMutation(source shared.OneCType) string {
	result := ""
	result += g.GenFilterStruct(source)
	result += "\n"
	result += g.GenFilterToStringFunc(source)
	result += "\n"
	result += g.GenFilterScalarToStringFunc(source)
	return result
}
