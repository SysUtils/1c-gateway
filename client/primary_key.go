package client

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenPrimaryKey(source shared.OneCType) string {
	result := g.GenPrimaryKeyStruct(source)
	result += "\n"
	result += g.GenPrimaryKeyTypeFunc(source)
	result += "\n"
	result += g.GenPrimaryKeySerializeFunc(source)
	return result
}

func (g *Generator) GenPrimaryKeyStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type Primary%s struct {\n", g.TranslateType(source.Name))
	for _, key := range source.Keys {
		for _, prop := range source.Properties {
			if prop.Name == key.Name {
				key.Type = prop.Type
				break
			}
		}
		result += "	"
		result += g.TranslateName(key.Name)
		result += " "
		result += g.TranslateType(key.Type)
		result += "\n"
	}

	return result + "}"
}

func (g *Generator) GenPrimaryKeyTypeFunc(source shared.OneCType) string {
	result := fmt.Sprintf("func (Primary%s) APIEntityType() string {\n", g.TranslateType(source.Name))
	result += fmt.Sprintf(`	return "%s"`+"\n", source.Name)
	return result + "}"
}

func (g *Generator) GenPrimaryKeySerializeFunc(source shared.OneCType) string {
	result := fmt.Sprintf("func (p Primary%s) Serialize() string {\n", g.TranslateType(source.Name))
	result += fmt.Sprintf(`	return `)
	for i, key := range source.Keys {
		if i > 0 {
			result += fmt.Sprintf(` + ",%s=" + `, key.Name)
		} else {
			result += fmt.Sprintf(`"%s=" + `, key.Name)
		}
		result += "p." + g.TranslateName(key.Name) + ".AsParameter()"
	}
	return result + "\n}"
}
