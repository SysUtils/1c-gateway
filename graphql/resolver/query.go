package resolver

import "gitlab.com/zullpro/core/1cclientgenerator.git/shared"

func (g *Generator) GenResolvers(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenEntityResolver(entity)
		result += "\n"
		result += g.GenEntitiesResolver(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenEntityResolver(source shared.OneCType) string {
	result := ""
	result += "func (r *GqlResolver) " + g.TranslateType(source.Name) + "(ctx context.Context, args " + g.TranslateType(source.Name) + "Args) (*" + g.TranslateType(source.Name) + ", error) {\n"
	result += "	return r.Client." + g.TranslateType(source.Name) + "(args.Key, nil)\n"
	result += "}"
	return result
}

func (g *Generator) GenEntitiesResolver(source shared.OneCType) string {
	result := ""
	result += "func (r *GqlResolver) " + g.TranslateType(source.Name) + "s(ctx context.Context, args " + g.TranslateType(source.Name) + "sArgs) (*[]" + g.TranslateType(source.Name) + ", error) {\n"
	result += "	if args.BaseWhere != nil {\n"
	result += "		if args.Filter != nil {\n"
	result += "			args.BaseWhere.Filter = *args.Filter\n"
	result += "		}\n"
	result += "		return r.Client." + g.TranslateType(source.Name) + "s(*args.BaseWhere)\n"
	result += "	}\n"
	result += "	if args.Filter != nil {\n"
	result += "		return r.Client." + g.TranslateType(source.Name) + "s(Where {Filter: *args.Filter})\n"
	result += "	}\n"
	result += "	return r.Client." + g.TranslateType(source.Name) + "s(Where {})\n"
	result += "}"
	return result
}
