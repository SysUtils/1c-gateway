package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genMutations(source []shared.OneCType, funcs []shared.Function) string {
	result := ""
	for _, entity := range source {
		result += g.genMutation(entity)
		result += "\n"
	}
	for _, f := range funcs {
		if f.IsSideEffecting {
			result += g.genMutationFunction(f)
			result += "\n"
		}

	}

	return result[:len(result)-1]
}

func (g *Generator) genMutation(source shared.OneCType) string {
	result := ""
	result += g.genMutationRemove(source)
	result += "\n"
	result += g.genMutationUpdate(source)
	result += "\n"
	result += g.genMutationCreate(source)
	return result
}

func (g *Generator) genMutationRemove(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf(
		`func (r *GqlResolver) Remove%s(ctx context.Context, args %sRemoveArgs) (bool, error) {
	err :=  r.Client.Delete%s(args.Key)
	return err == nil, err
}`, t, t, t)
	return result
}

func (g *Generator) genMutationFunction(source shared.Function) string {
	n := g.translateName(source.Name)
	if source.IsBindable {
		t := g.translateType(source.Parameters[0].Type)
		result := fmt.Sprintf(
			`func (r *GqlResolver) %s%s(ctx context.Context, args %s%sArgs) (bool, error) {
	_, err := r.Client.ExecuteEntityMethod(args.Key, args.Args);
	return err == nil, err
}`, n, t, n, t)
		return result
	}
	return ""
}

func (g *Generator) genMutationUpdate(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf(
		`func (r *GqlResolver) Update%s(ctx context.Context, args %sUpdateArgs) (*%s, error) {
	return r.Client.Update%s(args.Key, args.Entity)
}`, t, t, t, t)
	return result
}

func (g *Generator) genMutationCreate(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf(
		`func (r *GqlResolver) Create%s(ctx context.Context, args %sCreateArgs) (*%s, error) {
	return r.Client.Create%s(args.Entity)
}`, t, t, t, t)
	return result
}
