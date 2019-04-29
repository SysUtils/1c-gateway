package resolver

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
	"strings"
)

var ScalarTypes = map[string]bool{
	"String":   true,
	"Int":      true,
	"Int16":    true,
	"Int64":    true,
	"Double":   true,
	"Float":    true,
	"DateTime": true,
	"Binary":   true,
	"Stream":   true,
	"Boolean":  true,
	"Guid":     true,
}

func (g *Generator) GenFilters(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenFilter(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenFilter(source shared.OneCType) string {
	result := ""
	result += g.GenFilterStruct(source)
	result += "\n"
	result += g.GenFilterToStringFunc(source)
	result += "\n"
	result += g.GenFilterScalarToStringFunc(source)
	return result
}

func (g *Generator) GenFilterStruct(source shared.OneCType) string {
	result := ""
	result += "type " + g.TranslateType(source.Name) + "Filter struct {\n"
	result += "	AND *[]" + g.TranslateType(source.Name) + "Filter\n"
	result += "	OR *[]" + g.TranslateType(source.Name) + "Filter\n"
	for _, prop := range source.Properties {
		transType := g.TranslateType(prop.Type)
		if _, ok := ScalarTypes[transType]; ok {
			result += "	" + g.TranslateName(prop.Name) + "_eq *" + transType + "\n"
			result += "	" + g.TranslateName(prop.Name) + "_ne *" + transType + "\n"
		} else {
			if !strings.HasPrefix(transType, "[") {
				result += "	" + g.TranslateName(prop.Name) + ": " + transType + "Filter\n"
			}
		}
	}
	return result + "}"
}

func (g *Generator) GenFilterToStringFunc(source shared.OneCType) string {
	result := ""
	result += "func (t " + g.TranslateType(source.Name) + "Filter) ToString() string {\n"
	result += fmt.Sprintf(`	if t.AND != nil && len(*t.AND) > 0 {
		tmp := %sFilter{AND: t.AND}
		t.AND = nil
		*tmp.AND = append(*tmp.AND, t)
		result := "true"
		for _, f := range *tmp.AND {
			result += " and " + f.ToString()
		}
		return "("+result+")"
	}
	if t.OR != nil && len(*t.OR) > 0 {
		tmp := %sFilter{OR: t.OR}
		t.OR = nil
		*tmp.OR = append(*tmp.OR, t)
		result := "false"
		for _, f := range *tmp.OR {
			result += " or " + f.ToString()
		}
		return "("+result+")"
	}
	return t.ScalarToString()
}`, g.TranslateType(source.Name), g.TranslateType(source.Name))
	return result
}

func (g *Generator) GenFilterScalarToStringFunc(source shared.OneCType) string {
	result := ""
	result += "func (f *" + g.TranslateType(source.Name) + "Filter) ScalarToString() string {\n"
	result += `	result := "true"`
	result += "\n"
	for _, prop := range source.Properties {
		transType := g.TranslateType(prop.Type)
		if _, ok := ScalarTypes[transType]; ok {
			result += `	if f.` + g.TranslateName(prop.Name) + `_ne != nil {` + "\n"
			result += `		result += " and ` + prop.Name + ` ne " + f.` + g.TranslateName(prop.Name) + `_ne.AsParameter()` + "\n"
			result += "	}\n"
			result += `	if f.` + g.TranslateName(prop.Name) + `_eq != nil {` + "\n"
			result += `		result += " and ` + prop.Name + ` eq " + f.` + g.TranslateName(prop.Name) + `_eq.AsParameter()` + "\n"
			result += "	}\n"
		} else {
			if !strings.HasPrefix(transType, "[") {
				result += "	" + g.TranslateName(prop.Name) + ": " + transType + "Filter\n"
			}
		}
	}
	result += "	return `(` + result + `)`\n"
	return result + "}"
}
