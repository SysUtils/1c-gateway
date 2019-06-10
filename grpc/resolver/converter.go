package resolver

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genComplexConverters(source []shared.OneCType) string {
	result := ""
	for _, e := range source {
		result += g.genGrpcConverter(e)
		result += "\n"
		result += g.genNativeConverter(e)
		result += "\n"
	}
	return result
}

func (g *Generator) genTypeConverters(source []shared.OneCType) string {
	result := ""
	for _, e := range source {
		result += g.genGrpcConverter(e)
		result += "\n"
		result += g.genNativeConverter(e)
		result += "\n"
		result += g.genGrpcPrimaryConverter(e)
		result += "\n"
		result += g.genNativePrimaryConverter(e)
		result += "\n"
		result += g.genNativeFilterConverter(e)
		result += "\n"

	}
	return result
}

func (g *Generator) genGrpcPrimaryConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *Primary%s) ToGrpc() (%sPrimary, error) {\n", g.translateNativeType(source.Name), g.translateGrpcType(source.Name))
	result += fmt.Sprintf("	result := %sPrimary {}\n", g.translateGrpcType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genGrpcConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%s) ToGrpc() (%s, error) {\n", g.translateNativeType(source.Name), g.translateGrpcType(source.Name))
	result += fmt.Sprintf("	result := %s {}\n", g.translateGrpcType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genNativePrimaryConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%sPrimary) ToNative() (Primary%s, error) {\n", g.translateGrpcType(source.Name), g.translateNativeType(source.Name))
	result += fmt.Sprintf("	result := Primary%s {}\n", g.translateNativeType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genNativeFilterConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%sFilter) ToNative() (%sFilter, error) {\n", g.translateGrpcType(source.Name), g.translateNativeType(source.Name))
	result += fmt.Sprintf("	result := %sFilter {}\n", g.translateNativeType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}

func (g *Generator) genNativeConverter(source shared.OneCType) string {
	result := fmt.Sprintf("func (t *%s) ToNative() (%s, error) {\n", g.translateGrpcType(source.Name), g.translateNativeType(source.Name))
	result += fmt.Sprintf("	result := %s {}\n", g.translateNativeType(source.Name))
	result += fmt.Sprintf("	err := ConvertType(t, &result)\n")
	result += fmt.Sprintf("	return result, err\n")
	result += fmt.Sprintf("}")
	return result
}
