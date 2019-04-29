package resolver

import "gitlab.com/zullpro/core/1cclientgenerator.git/shared"

func (g *Generator) GenArgs(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenEntityArgs(entity)
		result += "\n"
		result += g.GenEntitiesArgs(entity)
		result += "\n"
		result += g.GenCreateArgs(entity)
		result += "\n"
		result += g.GenRemoveArgs(entity)
		result += "\n"
		result += g.GenUpdateArgs(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenEntityArgs(source shared.OneCType) string {
	result := ""
	result += "type " + g.TranslateType(source.Name) + "Args struct {\n"
	result += "	Key Primary" + g.TranslateType(source.Name)
	result += "\n}"
	return result
}

func (g *Generator) GenEntitiesArgs(source shared.OneCType) string {
	result := ""
	result += "type " + g.TranslateType(source.Name) + "sArgs struct {\n"
	result += "	BaseWhere *Where\n"
	result += "	Filter *" + g.TranslateType(source.Name) + "Filter\n"
	result += "}"
	return result
}

func (g *Generator) GenCreateArgs(source shared.OneCType) string {
	result := ""
	result += "type " + g.TranslateType(source.Name) + "CreateArgs struct {\n"
	result += "	Entity " + g.TranslateType(source.Name) + "\n"
	result += "}"
	return result
}

func (g *Generator) GenRemoveArgs(source shared.OneCType) string {
	result := ""
	result += "type " + g.TranslateType(source.Name) + "RemoveArgs struct {\n"
	result += "	Key Primary" + g.TranslateType(source.Name) + "\n"
	result += "}"
	return result
}

func (g *Generator) GenUpdateArgs(source shared.OneCType) string {
	result := ""
	result += "type " + g.TranslateType(source.Name) + "UpdateArgs struct {\n"
	result += "	Key Primary" + g.TranslateType(source.Name) + "\n"
	result += "	Entity " + g.TranslateType(source.Name) + "\n"
	result += "}"
	return result
}
