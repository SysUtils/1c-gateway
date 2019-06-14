package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genQueries(source []shared.OneCType) string {
	queries := "type Query {\n"
	for _, entity := range source {
		queries += g.genQuery(entity)
		queries += "\n"
	}
	queries += fmt.Sprintf(`}`)
	return queries
}

func (g *Generator) genQuery(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf(
		`	%s(Key: Primary%s!): %s
	%ss(Options: Options, Filter: %sFilter, OrderBy: %sField): [%s!]`, t, t, t, t, t, t, t)
	return result
}
