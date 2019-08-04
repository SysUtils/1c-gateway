package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genMutations(source []shared.OneCType, funcs []shared.Function) string {
	mutations := ""
	for _, entity := range source {
		mutations += g.genMutation(entity)
		mutations += "\n"
	}
	for _, f := range funcs {
		if f.IsSideEffecting {
			mutations += g.genMutationFunction(f)
			mutations += "\n"
		}

	}

	result := fmt.Sprintf(`type Mutation {
	%s}`, mutations)

	return result
}

func (g *Generator) genMutation(source shared.OneCType) string {
	t := g.translateType(source.Name)
	tInput := g.translateType(source.Name)
	result := fmt.Sprintf(
		`	create%s(entity: %sInput!): %s
	update%s(entity: %sInput!): %s
	remove%s(key: Primary%s!): Boolean!`, t, tInput, t, t, tInput, t, t, t)
	return result
}

func (g *Generator) genMutationFunction(source shared.Function) string {
	n := g.translateName(source.Name)
	if source.IsBindable {
		t := g.translateType(source.Parameters[0].Type)
		return fmt.Sprintf(
			`	%s%s(key: Primary%s!, args: %s%sArgs!): Boolean!`, n, t, t, n, t)
	}
	return ""
}
