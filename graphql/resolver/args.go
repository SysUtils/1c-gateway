package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genArgs(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genEntityArgs(entity)
		result += "\n"
		result += g.genEntitiesArgs(entity)
		result += "\n"
		result += g.genCreateArgs(entity)
		result += "\n"
		result += g.genRemoveArgs(entity)
		result += "\n"
		result += g.genUpdateArgs(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) genEntityArgs(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf("type %sArgs struct {\n", t)
	result += fmt.Sprintf("	Key Primary%s\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genEntitiesArgs(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf("type %ssArgs struct {\n", t)
	result += fmt.Sprintf("	BaseWhere *Where\n")
	result += fmt.Sprintf("	Filter *%sFilter\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genCreateArgs(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf("type %sCreateArgs struct {\n", t)
	result += fmt.Sprintf("	Entity %s\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genRemoveArgs(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf("type %sRemoveArgs struct {\n", t)
	result += fmt.Sprintf("	Key Primary%s\n", t)
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genUpdateArgs(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf("type %sUpdateArgs struct {\n", t)
	result += fmt.Sprintf("	Key Primary%s\n", t)
	result += fmt.Sprintf("	Entity %s\n", t)
	result += fmt.Sprintf("}")
	return result
}
