package native

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genFunctionsType(source []shared.Function) string {
	result := ""
	for _, entity := range source {
		result += g.genFunctionType(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) genFunctionType(source shared.Function) string {
	result := g.genFunctionStruct(source)
	result += "\n"
	result += g.genFunctionNameFunc(source)
	result += "\n"
	result += g.genFunctionParametersFunc(source)
	return result
}

func (g *Generator) genFunctionStruct(source shared.Function) string {
	result := ""
	if source.IsBindable {
		result += fmt.Sprintf("type Function%s%s struct {\n", g.translateType(source.Parameters[0].Type), g.translateName(source.Name))
	} else {
		result += fmt.Sprintf("type Function%s struct {\n", g.translateName(source.Name))
	}

	for i, p := range source.Parameters {
		if i == 0 && source.IsBindable {
			continue
		}
		result += "	"
		result += g.translateName(p.Name)
		result += " "
		result += g.translateType(p.Type)
		result += " `"
		result += fmt.Sprintf(`odata:"%s"`, p.Name)
		result += "`"
		result += "\n"
	}
	result += "}"

	return result
}

func (g *Generator) genFunctionNameFunc(source shared.Function) string {
	result := ""
	if source.IsBindable {
		result += fmt.Sprintf("func (Function%s%s) Name() string {\n", g.translateType(source.Parameters[0].Type), g.translateName(source.Name))
	} else {
		result += fmt.Sprintf("func (Function%s) Name() string {\n", g.translateName(source.Name))
	}
	result += fmt.Sprintf(`	return "%s"`, source.Name)
	result += "\n}"

	return result
}

func (g *Generator) genFunctionParametersFunc(source shared.Function) string {
	result := ""
	if source.IsBindable {
		result += fmt.Sprintf("func (f Function%s%s) Parameters() string {\n", g.translateType(source.Parameters[0].Type), g.translateName(source.Name))
	} else {
		result += fmt.Sprintf("func (f Function%s) Parameters() string {\n", g.translateName(source.Name))
	}
	result += fmt.Sprintf(`	return `)
	for i, key := range source.Parameters {
		if source.IsBindable && i == 0 {
			continue
		}
		if i == 0 || source.IsBindable && i == 1 {
			result += fmt.Sprintf(`"%s=" + `, key.Name)
		} else {
			result += fmt.Sprintf(` + "&%s=" + `, key.Name)

		}
		result += "f." + g.translateName(key.Name) + ".AsParameter()"
	}
	result += "\n}"

	return result
}
