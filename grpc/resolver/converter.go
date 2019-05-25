package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) GenComplexConverters(source []shared.OneCType) string {
	result := ""
	for _, e := range source {
		result += g.GenGrpcConverter(e)
		result += "\n"
		result += g.GenNativeConverter(e)
		result += "\n"
	}
	return result
}

func (g *Generator) GenTypeConverters(source []shared.OneCType) string {
	result := ""
	for _, e := range source {
		result += g.GenGrpcConverter(e)
		result += "\n"
		result += g.GenNativeConverter(e)
		result += "\n"
		result += g.GenGrpcPrimaryConverter(e)
		result += "\n"
		result += g.GenNativePrimaryConverter(e)
		result += "\n"
		result += g.GenNativeFilterConverter(e)
		result += "\n"

	}
	return result
}

func (g *Generator) GenGrpcPrimaryConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *Primary%s) ToGrpc() (%sPrimary, error) {\n", g.TranslateNativeType(source.Name), g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	result := %sPrimary {}\n", g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenGrpcConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%s) ToGrpc() (%s, error) {\n", g.TranslateNativeType(source.Name), g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	result := %s {}\n", g.TranslateGrpcType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenNativePrimaryConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%sPrimary) ToNative() (Primary%s, error) {\n", g.TranslateGrpcType(source.Name), g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	result := Primary%s {}\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenNativeFilterConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%sFilter) ToNative() (%sFilter, error) {\n", g.TranslateGrpcType(source.Name), g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	result := %sFilter {}\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) GenNativeConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%s) ToNative() (%s, error) {\n", g.TranslateGrpcType(source.Name), g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	result := %s {}\n", g.TranslateNativeType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}
