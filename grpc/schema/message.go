package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
)

func (g *Generator) genMessages(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genMessage(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) genMessage(source shared.OneCType) string {
	result := g.genTypeMessage(source)
	result += "\n"
	result += g.genQuerysResponseMessage(source)
	result += "\n"
	result += g.genQuerysWhereMessage(source)
	result += "\n"
	result += g.genFilterMessage(source)
	result += "\n"
	result += g.genRemoveResponseMessage(source)
	result += "\n"
	result += g.genPrimaryKeyMessage(source)
	result += "\n"
	result += g.genUpdateRequestMessage(source)
	return result
}

func (g *Generator) genTypeMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %s {\n", g.translateType(source.Name))
	for i, prop := range source.Properties {
		result += fmt.Sprintf("	%s %s = %d;\n", g.translateType(prop.Type), g.translateName(prop.Name), i+1)
	}
	return result + "}"
}

func (g *Generator) genQuerysResponseMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %ssResponse {\n", g.translateType(source.Name))
	result += fmt.Sprintf("	repeated %s Result = 1;", g.translateType(source.Name))
	return result + "}"
}

func (g *Generator) genQuerysWhereMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sWhere {\n", g.translateType(source.Name))
	result += fmt.Sprintf("	BaseWhere Base = 1;")
	result += fmt.Sprintf("	%sFilter Filter = 2;", g.translateType(source.Name))
	return result + "}"
}

func (g *Generator) genFilterMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sFilter {\n", g.translateType(source.Name))
	i := 1
	for _, prop := range source.Properties {
		propType := g.translateType(prop.Type)
		propName := g.translateName(prop.Name)
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

func (g *Generator) genPrimaryKeyMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sPrimary {\n", g.translateType(source.Name))
	for i, key := range source.Keys {
		for _, prop := range source.Properties {
			if prop.Name == key.Name {
				key.Type = prop.Type
				break
			}
		}
		result += fmt.Sprintf("	%s %s = %d;\n", g.translateType(key.Type), g.translateName(key.Name), i+1)
	}

	return result + "}"
}

func (g *Generator) genRemoveResponseMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message Remove%sResponse {\n", g.translateType(source.Name))
	result += "	bool Status = 1;\n"
	return result + "}"
}

func (g *Generator) genUpdateRequestMessage(source shared.OneCType) string {
	result := fmt.Sprintf("message %sUpdateRequest {\n", g.translateType(source.Name))
	result += fmt.Sprintf("	%sPrimary Key = 1;\n", g.translateType(source.Name))
	result += fmt.Sprintf("	%s Entity = 2;\n", g.translateType(source.Name))
	return result + "}"
}
