package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genFunctions(source []shared.Function) string {
	result := ""
	for _, f := range source {
		if f.IsSideEffecting && f.IsBindable {
			result += g.genFunction(f)
			result += "\n"
		}
	}
	return result[:len(result)-1]
}

func (g *Generator) genFunction(source shared.Function) string {
	result := g.genFunctionStruct(source)
	result += "\n"
	return result
}

func (g *Generator) genFunctionStruct(source shared.Function) string {
	result := ""
	if source.IsBindable {
		result += fmt.Sprintf("input %s%sArgs {\n", g.translateName(source.Name), g.translateType(source.Parameters[0].Type))
	}

	for i, p := range source.Parameters {
		if i == 0 && source.IsBindable {
			continue
		}
		result += "	"
		result += g.translateName(p.Name)
		result += ": "
		result += g.translateType(p.Type)
		result += "!\n"
	}
	result += "}"

	return result
}
