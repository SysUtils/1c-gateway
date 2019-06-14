package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genFields(source []shared.OneCType) string {
	queries := ""
	for _, entity := range source {
		queries += g.genField(entity)
		queries += "\n"
	}

	return queries
}

func (g *Generator) genField(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf("enum %sField {\n", t)
	for _, val := range source.Properties {
		translatedKey := g.translateName(val.Name)
		result += fmt.Sprintf(
			`	%s_desc,
	%s_asc,
`, translatedKey, translatedKey)
	}

	result += fmt.Sprintf(`}`)
	return result
}
