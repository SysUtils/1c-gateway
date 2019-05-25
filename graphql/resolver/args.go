package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

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
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf("type %sArgs struct {\n", t)
	result += fmt.Sprintf("	Key Primary%s\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenEntitiesArgs(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf("type %ssArgs struct {\n", t)
	result += fmt.Sprintf("	BaseWhere *Where\n")
	result += fmt.Sprintf("	Filter *%sFilter\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenCreateArgs(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf("type %sCreateArgs struct {\n", t)
	result += fmt.Sprintf("	Entity %s\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenRemoveArgs(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf("type %sRemoveArgs struct {\n", t)
	result += fmt.Sprintf("	Key Primary%s\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenUpdateArgs(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf("type %sUpdateArgs struct {\n", t)
	result += fmt.Sprintf("	Key Primary%s\n", t)
	result += fmt.Sprintf("	Entity %s\n", t)
	result += fmt.Sprintf("}")
	return result
}
