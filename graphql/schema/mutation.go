package schema

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenMutations(source []shared.OneCType) string {
	mutations := ""
	for _, entity := range source {
		mutations += g.GenMutation(entity)
		mutations += "\n"
	}

	result := fmt.Sprintf(`type Mutation {
	%s}`, mutations)

	return result
}

func (g *Generator) GenMutation(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	tInput := g.TranslateType(source.Name)
	result := fmt.Sprintf(
		`	Create%s(Entity: %s!): %s
	Update%s(Key: Primary%s!, Entity: %sInput!): %s\n"
	Remove%s(Key: Primary%s!): Boolean!`, t, tInput, t, t, t, tInput, t, t, t)
	return result
}
