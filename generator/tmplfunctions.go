package generator

import (
	"fmt"
	"log"
	"pkg.re/essentialkaos/translit.v2"
	"strings"
)

func Translit(text string) string {
	return translit.EncodeToICAO(text)
}

func BuildFilter(keys []PropertyRef) string {
	result := `[]string {`
	for _, key := range keys {
		switch key.Type {
		case "Int":
			result += fmt.Sprintf(`fmt.Sprintf("%s=%%d ", %s), `, key.Name, key.CamelName)
		case "Guid":
			result += fmt.Sprintf(`fmt.Sprintf("%s=guid'%%s' ", url.PathEscape(string(%s))), `, key.Name, key.CamelName)
		case "String":
			result += fmt.Sprintf(`fmt.Sprintf("%s='%%s' ", url.PathEscape(string(%s))), `, key.Name, key.CamelName)
		case "Float":
			result += fmt.Sprintf(`fmt.Sprintf("%s=%%f ", %s), `, key.Name, key.CamelName)
		default:
			log.Println("Wrong type for key: ", key.Type)
		}
	}
	return result + `}`
}

func BuildEntityRefkeys(keys []PropertyRef) string {
	result := `[]string {`
	for _, key := range keys {
		switch key.Type {
		case "Int":
			result += fmt.Sprintf(`fmt.Sprintf("%s=%%d ", entity.%s), `, key.Name, Translit(key.CamelName))
		case "Guid":
			result += fmt.Sprintf(`fmt.Sprintf("%s=guid'%%s' ", url.PathEscape(string(entity.%s))), `, key.Name, Translit(key.CamelName))
		case "String":
			result += fmt.Sprintf(`fmt.Sprintf("%s='%%s' ", url.PathEscape(string(entity.%s))), `, key.Name, Translit(key.CamelName))
		case "Float":
			result += fmt.Sprintf(`fmt.Sprintf("%s=%%f ", entity.%s), `, key.Name, Translit(key.CamelName))
		default:
			log.Println("Wrong type for key: ", key.Type)
		}
	}
	return result + `}`
}

func BuildArgs(keys []Parameter) string {
	result := `map[string]string { `
	for _, key := range keys {
		switch key.Type {
		case "Int":
			result += fmt.Sprintf(`"%s": fmt.Sprintf("%%d", %s), `, key.Name, key.Name)
		case "Guid":
			result += fmt.Sprintf(`"%s": fmt.Sprintf("guid'%%s'", %s), `, key.Name, key.Name)
		case "String":
			result += fmt.Sprintf(`"%s": fmt.Sprintf("'%%s'", url.PathEscape(string(%s))), `, key.Name, key.Name)
		case "Float":
			result += fmt.Sprintf(`"%s": fmt.Sprintf("%%f ", %s), `, key.Name, key.Name)
		default:
			log.Println("Wrong type for key: ", key.Type)
		}
	}
	return result + `}`
}

func ConvertToGqlType(text string) string {
	for strings.HasPrefix(text, "[]") {
		return fmt.Sprintf("[%s!]", ConvertToGqlType(text[2:]))
	}
	return text
}
