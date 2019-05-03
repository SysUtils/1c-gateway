package native

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenFunctions(source []shared.Function) string {
	result := ""
	for _, entity := range source {
		result += g.GenFunction(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenFunction(source shared.Function) string {
	result := g.GenFunctionStruct(source)
	result += "\n"
	result += g.GenFunctionNameFunc(source)
	result += "\n"
	result += g.GenFunctionParametersFunc(source)
	return result
}

func (g *Generator) GenFunctionStruct(source shared.Function) string {
	result := ""
	if source.IsBindable {
		result += fmt.Sprintf("type Function%s%s struct {\n", g.TranslateType(source.Parameters[0].Type), g.TranslateName(source.Name))
	} else {
		result += fmt.Sprintf("type Function%s struct {\n", g.TranslateName(source.Name))
	}

	for i, p := range source.Parameters {
		if i == 0 && source.IsBindable {
			continue
		}
		result += "	"
		result += g.TranslateName(p.Name)
		result += " "
		result += g.TranslateType(p.Type)
		result += " `"
		result += fmt.Sprintf(`odata:"%s"`, p.Name)
		result += "`"
		result += "\n"
	}
	result += "}"

	return result
}

func (g *Generator) GenFunctionNameFunc(source shared.Function) string {
	result := ""
	if source.IsBindable {
		result += fmt.Sprintf("func (Function%s%s) Name() string {\n", g.TranslateType(source.Parameters[0].Type), g.TranslateName(source.Name))
	} else {
		result += fmt.Sprintf("func (Function%s) Name() string {\n", g.TranslateName(source.Name))
	}
	result += fmt.Sprintf(`	return "%s"`, source.Name)
	result += "\n}"

	return result
}

func (g *Generator) GenFunctionParametersFunc(source shared.Function) string {
	result := ""
	if source.IsBindable {
		result += fmt.Sprintf("func (f Function%s%s) Parameters() string {\n", g.TranslateType(source.Parameters[0].Type), g.TranslateName(source.Name))
	} else {
		result += fmt.Sprintf("func (f Function%s) Parameters() string {\n", g.TranslateName(source.Name))
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
		result += "f." + g.TranslateName(key.Name) + ".AsParameter()"
	}
	result += "\n}"

	return result
}
