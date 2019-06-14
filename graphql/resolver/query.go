package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
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

	fieldmapBody := ""
	for _, prop := range source.Properties {
		key := strings.Replace(g.translateName(prop.Name), `"`, `\"`, -1)
		val := strings.Replace(prop.Name, `"`, `\"`, -1)
		fieldmapBody += fmt.Sprintf(`    "%s_asc": "%s asc",
    "%s_desc": "%s desc",
`, key, val, key, val)
	}

	result := fmt.Sprintf(
		`
var %sFieldMap = map[string]string {
%s}
func (r *GqlResolver) %ss(ctx context.Context, args %ssArgs) (*[]%s, error) {
	where := Where {}
	if args.Options != nil {
		where = *args.Options
	}
	if args.Filter != nil {
		where.Filter = *args.Filter
	}
	if args.OrderBy != nil {
		where.Orderby = %sFieldMap[*args.OrderBy]
	}
	return r.Client.%ss(where)
}`, t, fieldmapBody, t, t, t, t, t)

	return result
}
