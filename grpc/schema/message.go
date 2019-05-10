package schema

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenMessages(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenMessage(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) GenMessage(source shared.OneCType) string {
	result := g.GenTypeMessage(source)
	result += "\n"
	result += g.GenPrimaryKeyMessage(source)
	return result
}

func (g *Generator) GenTypeMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %s {\n", g.TranslateType(source.Name))
	for i, prop := range source.Properties {
		result += fmt.Sprintf("	%s %s = %d;\n", g.TranslateType(prop.Type), g.TranslateName(prop.Name), i+1)
	}
	return result + "}"
}

func (g *Generator) GenPrimaryKeyMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sPrimary {\n", g.TranslateType(source.Name))
	for i, key := range source.Keys {
		for _, prop := range source.Properties {
			if prop.Name == key.Name {
				key.Type = prop.Type
				break
			}
		}
		result += fmt.Sprintf("	%s %s = %d;\n", g.TranslateType(key.Type), g.TranslateName(key.Name), i+1)
	}

	return result + "}"
}
