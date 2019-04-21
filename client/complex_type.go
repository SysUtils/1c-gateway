package client

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenComplexType(source shared.OneCType) string {
	result := g.GenComplexTypeStruct(source)
	return result
}

func (g *Generator) GenComplexTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type %s struct {\n", g.TranslateType(source.Name))
	for _, prop := range source.Properties {
		result += "	"
		result += g.TranslateName(prop.Name)
		result += " "
		if prop.Nullable {
			result += "*"
		}
		result += g.TranslateType(prop.Type)
		result += " `"
		result += fmt.Sprintf(`json:"%s,omitempty"`, prop.Name)
		result += "`"
		result += "\n"
	}
	return result + "}"
}
