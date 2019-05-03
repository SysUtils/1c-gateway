package schema

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
	"log"
)

func (g *Generator) ExtractAssociations(source []shared.Association) {
	for _, assoc := range source {
		name := "StandardODATA." + assoc.Name
		if _, ok := g.Associations[name]; !ok {
			g.Associations[name] = make(map[string]string, len(assoc.Ends))
		}
		for _, end := range assoc.Ends {
			g.Associations[name][end.Role] = end.Type
		}
	}
}

func (g *Generator) GenTypes(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenType(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) GenType(source shared.OneCType) string {
	result := g.GenTypeStruct(source)
	return result
}

func (g *Generator) GenTypeStruct(source shared.OneCType) string {
	result := g.GenInputTypeStruct(source)
	result += "\n"
	result += g.GenOutputTypeStruct(source)

	return result
}

func (g *Generator) GenInputTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("input %sInput {\n", g.TranslateType(source.Name))
	for _, prop := range source.Properties {
		result += fmt.Sprintf("	%s: %s", g.TranslateName(prop.Name), g.TranslateInputType(prop.Type))
		if !prop.Nullable {
			result += "!"
		}
		result += "\n"
	}
	return result + "}"
}

func (g *Generator) GenOutputTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("input %s {\n", g.TranslateType(source.Name))
	for _, prop := range source.Properties {
		result += fmt.Sprintf("	%s: %s", g.TranslateName(prop.Name), g.TranslateInputType(prop.Type))
		if !prop.Nullable {
			result += "!"
		}
		result += "\n"
	}
	for _, nav := range source.Navigations {
		if _, ok := g.Associations[nav.Type]; !ok {
			log.Panicf("navigation not found: %s", nav.Type)
		}
		if _, ok := g.Associations[nav.Type][nav.ToRole]; !ok {
			log.Panicf("navigation role not found: %s.%s", nav.Type, nav.ToRole)
		}
		result += fmt.Sprintf("	%s: %s\n", g.TranslateName(nav.Name), g.TranslateType(g.Associations[nav.Type][nav.ToRole]))
	}
	return result + "}"
}
