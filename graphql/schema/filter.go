package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
)

func (g *Generator) GenFilters(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenFilter(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) GenFilter(source shared.OneCType) string {
	name := g.TranslateType(source.Name) + "Filter"
	result := fmt.Sprintf(`input %s {
	AND: [%s!]
	OR: [%s!]
`, name, name, name)
	for _, prop := range source.Properties {
		propType := g.TranslateType(prop.Type)
		propName := g.TranslateName(prop.Name)
		if _, ok := ScalarTypes[propType]; ok {
			result += fmt.Sprintf(
				`	%sEq: %s
	%sNe: %s
`, propName, propType, propName, propType)
		} else {
			if !strings.HasPrefix(propType, "[") {
				result += fmt.Sprintf("	%s: %sFilter\n", propName, propType)
			}
		}
	}
	result += "}"
	return result
}
