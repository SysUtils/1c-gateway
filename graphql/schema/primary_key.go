package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genPrimaryKeys(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genPrimaryKey(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) genPrimaryKey(source shared.OneCType) string {
	result := g.genPrimaryKeyStruct(source)
	return result
}

func (g *Generator) genPrimaryKeyStruct(source shared.OneCType) string {
	result := fmt.Sprintf("input Primary%s {\n", g.translateType(source.Name))
	for _, key := range source.Keys {
		for _, prop := range source.Properties {
			if prop.Name == key.Name {
				key.Type = prop.Type
				break
			}
		}
		result += fmt.Sprintf("	%s: %s!\n", g.translateName(key.Name), g.translateType(key.Type))
	}
	result += fmt.Sprintf("}")
	return result
}
