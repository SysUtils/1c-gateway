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
		result += g.GenQuerysResolver(e)
		result += "\n"
		result += g.GenRemoveResolver(e)
		result += "\n"
		result += g.GenUpdateResolver(e)
		result += "\n"
		result += g.GenCreateResolver(e)
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

func (g *Generator) GenQuerysResolver(source shared.OneCType) string {
	result := fmt.Sprintf("func (r *GrpcResolver)%ss(ctx context.Context, in *%sWhere) (*%ssResponse, error) {\n", g.TranslateNativeType(source.Name), g.TranslateGrpcType(source.Name), g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	where := in.Base.ToNative()\n")
	result += fmt.Sprintf("	filter, err := in.Filter.ToNative()\n")
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	where.Filter = filter\n")

	result += fmt.Sprintf("	res := %ssResponse{}\n", g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	native, err := r.Client.%ss(where)\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	err = ConvertType(native, &res.Result)\n")
	result += fmt.Sprintf("	return &res, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenRemoveResolver(source shared.OneCType) string {
	result := fmt.Sprintf("func (r *GrpcResolver)Remove%s(ctx context.Context, in *%sPrimary) (*Remove%sResponse, error) {\n", g.TranslateNativeType(source.Name), g.TranslateGrpcType(source.Name), g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	key, err := in.ToNative()\n")
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	err = r.Client.Delete%s(key)\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	return &Remove%sResponse{Status: (err!=nil)}, err\n", g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenUpdateResolver(source shared.OneCType) string {
	result := fmt.Sprintf("func (r *GrpcResolver)Update%s(ctx context.Context, in *%sUpdateRequest) (*%s, error) {\n", g.TranslateNativeType(source.Name), g.TranslateGrpcType(source.Name), g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	key, err := in.Key.ToNative()\n")
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	entity, err := in.Entity.ToNative()\n")
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	native, err := r.Client.Update%s(key, entity)\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	res, err := native.ToGrpc()\n")
	result += fmt.Sprintf("	return &res, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenCreateResolver(source shared.OneCType) string {
	result := fmt.Sprintf("func (r *GrpcResolver)Create%s(ctx context.Context, in *%s) (*%s, error) {\n", g.TranslateNativeType(source.Name), g.TranslateGrpcType(source.Name), g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	entity, err := in.ToNative()\n")
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	native, err := r.Client.Create%s(entity)\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	if err != nil {\n")
	result += fmt.Sprintf("		return nil, err\n")
	result += fmt.Sprintf("	}\n")
	result += fmt.Sprintf("	res, err := native.ToGrpc()\n")
	result += fmt.Sprintf("	return &res, err\n")
	result += fmt.Sprintf("}")
	return result
}
