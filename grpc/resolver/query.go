package resolver

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenQueryResolvers(source []shared.OneCType) string {
	result := ""
	for _, e := range source {
		result += g.GenQueryResolver(e)
		result += "\n"
	}
	return result
}

func (g *Generator) GenQueryResolver(source shared.OneCType) string {
	result := fmt.Sprintf("func (r *GrpcResolver)%s(ctx context.Context, in *%sPrimary) (*%s, error) {\n", g.TranslateNativeType(source.Name), g.TranslateGrpcType(source.Name), g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	key, err := in.ToNative()\n")
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	native, err := r.Client.%s(key, nil)\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	res, err := native.ToGrpc()\n")
	result += fmt.Sprintf("	return &res, err\n")
	result += fmt.Sprintf("}")
	return result
}
