package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genFilters(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genFilter(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) genFilter(source shared.OneCType) string {
	name := g.translateType(source.Name) + "Filter"
	result := fmt.Sprintf(`input %s {
	AND: [%s!]
	OR: [%s!]
`, name, name, name)
	for _, prop := range source.Properties {
		propType := g.translateType(prop.Type)
		propName := g.translateName(prop.Name)
		if _, ok := ScalarTypes[propType]; ok {
			if prop.Type == "Boolean" {
				result += fmt.Sprintf(
					`	%sEq: %s
	%sNe: %s
`, propName, propType, propName, propType)
			} else {
				result += fmt.Sprintf(
					`	%sEq: %s
	%sNe: %s
	%sGt: %s
	%sLt: %s
	%sGe: %s
	%sLe: %s
`, propName, propType, propName, propType, propName, propType, propName, propType, propName, propType, propName, propType)
			}
		} else {
			//if !strings.HasPrefix(propType, "[") {
			//	result += fmt.Sprintf("	%s: %sFilter\n", propName, propType)
			//}
		}
	}
	result += "}"
	return result
}
