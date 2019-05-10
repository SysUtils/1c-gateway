package schema

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
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
	return result
}
