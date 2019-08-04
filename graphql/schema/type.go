package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"log"
)

func (g *Generator) extractAssociations(source []shared.Association) {
	for _, assoc := range source {
		name := "StandardODATA." + assoc.Name
		if _, ok := g.Associations[name]; !ok {
			g.Associations[name] = make(map[string]string, len(assoc.Ends))
		}
		for _, end := range assoc.Ends {
			g.Associations[name][end.Role] = end.Type
		}
	}
}

func (g *Generator) genTypes(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genType(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) genType(source shared.OneCType) string {
	result := g.genTypeStruct(source)
	return result
}

func (g *Generator) genTypeStruct(source shared.OneCType) string {
	result := g.genInputTypeStruct(source)
	result += "\n"
	result += g.genOutputTypeStruct(source)

	return result
}

func (g *Generator) genInputTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("input %sInput {\n", g.translateType(source.Name))
	for _, prop := range source.Properties {
		result += fmt.Sprintf("	%s: %s", g.translateName(prop.Name), g.translateInputType(prop.Type))
		//if !prop.Nullable {
		//	result += "!"
		//}
		result += "\n"
	}
	return result + "}"
}

func (g *Generator) genOutputTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type %s {\n", g.translateType(source.Name))
	for _, prop := range source.Properties {
		result += fmt.Sprintf("	%s: %s", g.translateName(prop.Name), g.translateType(prop.Type))
		//if !prop.Nullable {
		//	result += "!"
		//}
		result += "\n"
	}
	for _, nav := range source.Navigations {
		if _, ok := g.Associations[nav.Type]; !ok {
			log.Panicf("navigation not found: %s", nav.Type)
		}
		if _, ok := g.Associations[nav.Type][nav.ToRole]; !ok {
			log.Panicf("navigation role not found: %s.%s", nav.Type, nav.ToRole)
		}
		result += fmt.Sprintf("	%s: %s\n", g.translateName(nav.Name), g.translateType(g.Associations[nav.Type][nav.ToRole]))
	}
	return result + "}"
}
