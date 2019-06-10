package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genService(source []shared.OneCType) string {
	result := "service GrpcOdata {\n"
	for _, entity := range source {
		result += g.genServiceMethod(entity)
		result += "\n"
	}
	return result + "}"
}

func (g *Generator) genServiceMethod(source shared.OneCType) string {
	result := fmt.Sprintf("	rpc %s(%sPrimary) returns (%s);", g.translateNativeType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc %ss(%sWhere) returns (%ssResponse);", g.translateNativeType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc Remove%s(%sPrimary) returns (Remove%sResponse);", g.translateNativeType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc Update%s(%sUpdateRequest) returns (%s);", g.translateNativeType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	result += "\n"
	result += fmt.Sprintf("	rpc Create%s(%s) returns (%s);", g.translateNativeType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	return result
}
