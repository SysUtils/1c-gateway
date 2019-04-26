package schema

import (
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

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
	return result + "}"
}
