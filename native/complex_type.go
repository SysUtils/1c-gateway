package native

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genComplexTypes(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genComplexType(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) genComplexType(source shared.OneCType) string {
	result := g.genComplexTypeStruct(source)
	return result
}

func (g *Generator) genComplexTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type %s struct {\n", g.translateType(source.Name))
	for _, prop := range source.Properties {
		result += "	"
		result += g.translateName(prop.Name)
		result += " "
		//if prop.Nullable {
		result += "*"
		//}
		result += g.translateType(prop.Type)
		result += " `"
		result += fmt.Sprintf(`json:"%s,omitempty"`, prop.Name)
		result += "`"
		result += "\n"
	}
	return result + "}"
}
