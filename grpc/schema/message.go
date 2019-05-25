package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
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
	result += g.GenQuerysResponseMessage(source)
	result += "\n"
	result += g.GenQuerysWhereMessage(source)
	result += "\n"
	result += g.GenFilterMessage(source)
	result += "\n"
	result += g.GenRemoveResponseMessage(source)
	result += "\n"
	result += g.GenPrimaryKeyMessage(source)
	result += "\n"
	result += g.GenUpdateRequestMessage(source)
	return result
}

func (g *Generator) GenTypeMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %s {\n", g.TranslateType(source.Name))
	for i, prop := range source.Properties {
		result += fmt.Sprintf("	%s %s = %d;\n", g.TranslateType(prop.Type), g.TranslateName(prop.Name), i+1)
	}
	return result + "}"
}

func (g *Generator) GenQuerysResponseMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %ssResponse {\n", g.TranslateType(source.Name))
	result += fmt.Sprintf("	repeated %s Result = 1;", g.TranslateType(source.Name))
	return result + "}"
}

func (g *Generator) GenQuerysWhereMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sWhere {\n", g.TranslateType(source.Name))
	result += fmt.Sprintf("	BaseWhere Base = 1;")
	result += fmt.Sprintf("	%sFilter Filter = 2;", g.TranslateType(source.Name))
	return result + "}"
}

func (g *Generator) GenFilterMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sFilter {\n", g.TranslateType(source.Name))
	i := 1
	for _, prop := range source.Properties {
		propType := g.TranslateType(prop.Type)
		propName := g.TranslateName(prop.Name)
		if _, ok := ScalarTypes[propType]; ok {
			result += fmt.Sprintf(
				`	%s %s_eq = %d;
	%s %s_ne = %d;
`, propType, propName, i, propType, propName, i+1)
			i += 2
		} else {
			if !strings.HasPrefix(propType, "repeated ") {
				result += fmt.Sprintf("	%sFilter %s = %d;\n", propType, propName, i)
				i += 1
			}
		}
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

func (g *Generator) GenRemoveResponseMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message Remove%sResponse {\n", g.TranslateType(source.Name))
	result += "	bool Status = 1;\n"
	return result + "}"
}

func (g *Generator) GenUpdateRequestMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sUpdateRequest {\n", g.TranslateType(source.Name))
	result += fmt.Sprintf("	%sPrimary Key = 1;\n", g.TranslateType(source.Name))
	result += fmt.Sprintf("	%s Entity = 2;\n", g.TranslateType(source.Name))
	return result + "}"
}
