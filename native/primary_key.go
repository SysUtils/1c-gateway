package native

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
	result += "\n"
	result += g.genPrimaryKeyTypeFunc(source)
	result += "\n"
	result += g.genPrimaryKeySerializeFunc(source)
	return result
}

func (g *Generator) genPrimaryKeyStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type Primary%s struct {\n", g.translateType(source.Name))
	for _, key := range source.Keys {
		for _, prop := range source.Properties {
			if prop.Name == key.Name {
				key.Type = prop.Type
				break
			}
		}
		result += "	"
		result += g.translateName(key.Name)
		result += " "
		result += g.translateType(key.Type)
		result += " `"
		result += fmt.Sprintf(`json:"%s,omitempty"`, key.Name)
		result += "`"
		result += "\n"
	}

	return result + "}"
}

func (g *Generator) genPrimaryKeyTypeFunc(source shared.OneCType) string {
	result := fmt.Sprintf("func (Primary%s) APIEntityType() string {\n", g.translateType(source.Name))
	result += fmt.Sprintf(`	return "%s"`+"\n", source.Name)
	return result + "}"
}

func (g *Generator) genPrimaryKeySerializeFunc(source shared.OneCType) string {
	result := fmt.Sprintf("func (p Primary%s) Serialize() string {\n", g.translateType(source.Name))
	result += fmt.Sprintf(`	return `)
	for i, key := range source.Keys {
		if i > 0 {
			result += fmt.Sprintf(` + ",%s=" + `, key.Name)
		} else {
			result += fmt.Sprintf(`"%s=" + `, key.Name)
		}
		result += "p." + g.translateName(key.Name) + ".AsParameter()"
	}
	return result + "\n}"
}
