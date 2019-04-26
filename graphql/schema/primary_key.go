package schema

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenPrimaryKeys(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenPrimaryKey(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenPrimaryKey(source shared.OneCType) string {
	result := g.GenPrimaryKeyStruct(source)
	return result
}

func (g *Generator) GenPrimaryKeyStruct(source shared.OneCType) string {
	result := fmt.Sprintf("input Primary%s {\n", g.TranslateType(source.Name))
	for _, key := range source.Keys {
		for _, prop := range source.Properties {
			if prop.Name == key.Name {
				key.Type = prop.Type
				break
			}
		}
		result += "	"
		result += g.TranslateName(key.Name)
		result += ": "
		result += g.TranslateType(key.Type)
		result += "!\n"
	}

	return result + "}"
}
