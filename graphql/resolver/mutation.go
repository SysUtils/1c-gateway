package resolver

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenMutations(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenMutation(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenMutation(source shared.OneCType) string {
	result := ""
	result += g.GenMutationRemove(source)
	result += "\n"
	result += g.GenMutationUpdate(source)
	result += "\n"
	result += g.GenMutationCreate(source)
	return result
}

func (g *Generator) GenMutationRemove(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf(
		`func (r *GqlResolver) Remove%s(ctx context.Context, args %sRemoveArgs) (bool, error) {
	err :=  r.Client.Delete%s(args.Key)
	return err == nil, err
}`, t, t, t)
	return result
}

func (g *Generator) GenMutationUpdate(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf(
		`func (r *GqlResolver) Update%s(ctx context.Context, args %sUpdateArgs) (*%s, error) {
	return r.Client.Update%s(args.Key, args.Entity)
}`, t, t, t, t)
	return result
}

func (g *Generator) GenMutationCreate(source shared.OneCType) string {
	t := g.TranslateType(source.Name)
	result := fmt.Sprintf(
		`func (r *GqlResolver) Create%s(ctx context.Context, args %sCreateArgs) (*%s, error) {
	return r.Client.Create%s(args.Entity)
}`, t, t, t, t)
	return result
}
