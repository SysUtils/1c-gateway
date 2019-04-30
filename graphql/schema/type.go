package schema

import (
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

func (g *Generator) GenType(source shared.OneCType) string {
	result := g.GenTypeStruct(source)
	result += "\n"
	return result
}

func (g *Generator) GenTypeStruct(source shared.OneCType) string {
	result := g.GenInputTypeStruct(source)
	result += "\n"
	result += g.GenOutputTypeStruct(source)

	return result
}

func (g *Generator) GenInputTypeStruct(source shared.OneCType) string {
	result := "input " + g.TranslateType(source.Name) + "Input {\n"
	for _, prop := range source.Properties {
		result += "	"
		result += g.TranslateName(prop.Name)
		result += ": "
		result += g.TranslateInputType(prop.Type)
		if !prop.Nullable {
			result += "!"
		}
		result += "\n"
	}
	return result + "}"
}

func (g *Generator) GenOutputTypeStruct(source shared.OneCType) string {
	result := "type " + g.TranslateType(source.Name) + " {\n"
	for _, prop := range source.Properties {
		result += "	"
		result += g.TranslateName(prop.Name)
		result += ": "
		result += g.TranslateType(prop.Type)
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
		result += "	"
		result += g.TranslateName(nav.Name)
		result += ": "
		result += g.TranslateType(g.Associations[nav.Type][nav.ToRole])
		result += "\n"
	}
	return result + "}"
}
