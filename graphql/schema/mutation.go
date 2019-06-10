package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genMutations(source []shared.OneCType) string {
	mutations := ""
	for _, entity := range source {
		mutations += g.genMutation(entity)
		mutations += "\n"
	}

	result := fmt.Sprintf(`type Mutation {
	%s}`, mutations)

	return result
}

func (g *Generator) genMutation(source shared.OneCType) string {
	t := g.translateType(source.Name)
	tInput := g.translateType(source.Name)
	result := fmt.Sprintf(
		`	Create%s(Entity: %sInput!): %s
	Update%s(Key: Primary%s!, Entity: %sInput!): %s
	Remove%s(Key: Primary%s!): Boolean!`, t, tInput, t, t, t, tInput, t, t, t)
	return result
}
