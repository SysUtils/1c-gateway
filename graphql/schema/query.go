package schema

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenQueries(source []shared.OneCType) string {
	queries := "type Query {\n"
	for _, entity := range source {
		queries += g.GenQuery(entity)
		queries += "\n"
	}
	queries += fmt.Sprintf(`}`)
	return queries
}

func (g *Generator) GenQuery(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf(
		`	%s(Key: Primary%s!): %s
	%ss(BaseWhere: BaseWhere, Filter: %sFilter): [%s!]`, t, t, t, t, t, t)
	return result
}
