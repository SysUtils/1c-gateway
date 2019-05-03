package resolver

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenConverters(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenConverterToGrpc(entity)
		result += "\n"
		//result += g.GenConverterFromGrpc(entity)
		//result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenConverterToGrpc(source shared.OneCType) string {
	result := fmt.Sprintf("func (e *%s) ToGrpc() %sGrpc {", g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += fmt.Sprintf("	return %sGrpc {", g.TranslateType(source.Name))
	for _, prop := range source.Properties {
		if !prop.Nullable {
			result += fmt.Sprintf("%s: e.%s,", g.TranslateName(prop.Name), g.TranslateName(prop.Name))
		}
	}
	result += "\n	}\n}"
	return result
}
