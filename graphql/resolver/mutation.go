package resolver

import "gitlab.com/zullpro/core/1cclientgenerator.git/shared"

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
	return result
}

func (g *Generator) GenMutationRemove(source shared.OneCType) string {
	result := "func (r *GqlResolver) Remove" + g.TranslateType(source.Name) + "(ctx context.Context, args " + g.TranslateType(source.Name) + "RemoveArgs) (bool, error) {\n"
	result += "	err :=  r.Client.Delete" + g.TranslateType(source.Name) + "(args.Key)\n"
	result += "	return err == nil, err\n"
	result += "}"
	return result
}
func (g *Generator) GenMutationUpdate(source shared.OneCType) string {
	result := "func (r *GqlResolver) Update" + g.TranslateType(source.Name) + "(ctx context.Context, args " + g.TranslateType(source.Name) + "UpdateArgs) (*" + g.TranslateType(source.Name) + ", error) {\n"
	result += "	return r.Client.Update" + g.TranslateType(source.Name) + "(args.Key, args.Entity)\n"
	result += "}"
	return result
}
