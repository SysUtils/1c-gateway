package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) GenService(source []shared.OneCType) string {
	result := "service GrpcOdata {\n"
	for _, entity := range source {
		result += g.GenServiceMethod(entity)
		result += "\n"
	}
	return result + "}"
}

func (g *Generator) GenServiceMethod(source shared.OneCType) string {
	result := fmt.Sprintf("	rpc %s(%sPrimary) returns (%s);", g.TranslateNativeType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc %ss(%sWhere) returns (%ssResponse);", g.TranslateNativeType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc Remove%s(%sPrimary) returns (Remove%sResponse);", g.TranslateNativeType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc Update%s(%sUpdateRequest) returns (%s);", g.TranslateNativeType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc Create%s(%s) returns (%s);", g.TranslateNativeType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	return result
}
