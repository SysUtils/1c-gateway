package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genResolvers(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genEntityResolver(entity)
		result += "\n"
		result += g.genEntitiesResolver(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) genEntityResolver(source shared.OneCType) string {
	t := g.translateType(source.Name)
	result := fmt.Sprintf(
		`func (r *GqlResolver) %s(ctx context.Context, args %sArgs) (*%s, error) {
	return r.Client.%s(args.Key, nil)
}`, t, t, t, t)
	return result
}

func (g *Generator) genEntitiesResolver(source shared.OneCType) string {
	t := g.translateType(source.Name)

	result := fmt.Sprintf(
		`func (r *GqlResolver) %ss(ctx context.Context, args %ssArgs) (*[]%s, error) {
	if args.BaseWhere != nil {
		if args.Filter != nil {
			args.BaseWhere.Filter = *args.Filter
		}
		return r.Client.%ss(*args.BaseWhere)
	}
	if args.Filter != nil {
		return r.Client.%ss(Where {Filter: *args.Filter})
	}
	return r.Client.%ss(Where {})
}`, t, t, t, t, t, t)

	return result
}
