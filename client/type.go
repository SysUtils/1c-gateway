package client

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/shared"
)

func (g *Generator) GenType(source shared.OneCType) string {
	result := g.GenTypeStruct(source)
	result += "\n"
	result += g.GenTypeFunc(source)
	result += "\n"
	result += g.GenNew(source)
	result += "\n"
	result += g.GenCreate(source)
	result += "\n"
	result += g.GenRead(source)
	result += "\n"
	result += g.GenUpdate(source)
	result += "\n"
	result += g.GenDelete(source)
	return result
}

func (g *Generator) GenTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type %s struct {\n", g.TranslateType(source.Name))
	for _, prop := range source.Properties {
		result += "	"
		result += g.TranslateName(prop.Name)
		result += " "
		if prop.Nullable {
			result += "*"
		}
		result += g.TranslateType(prop.Type)
		result += " `"
		result += fmt.Sprintf(`json:"%s,omitempty"`, prop.Name)
		result += "`"
		result += "\n"
	}
	return result + "}"
}

func (g *Generator) GenNew(source shared.OneCType) string {
	result := fmt.Sprintf("func New%s(data string, prevError error) (*%s, error) {\n", g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += fmt.Sprintf(`	if prevError != nil {
		return nil, prevError
	}
	result := new(%s)
	err := json.Unmarshal([]byte(data), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}`, g.TranslateType(source.Name))
	return result
}

func (g *Generator) GenTypeFunc(source shared.OneCType) string {
	result := fmt.Sprintf("func (%s) APIEntityType() string {\n", g.TranslateType(source.Name))
	result += fmt.Sprintf(`	return "%s"
}`, source.Name)
	return result
}

func (g *Generator) GenCreate(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) Create%s(entity %s) (*%s, error) {\n", g.TranslateType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += fmt.Sprintf(`	return New%s(c.createEntity(entity))`+"\n", g.TranslateType(source.Name))
	return result + "}"
}

func (g *Generator) GenRead(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) %s(key Primary%s, fields []string) (*%s, error) {\n", g.TranslateType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += fmt.Sprintf(`	return New%s(c.getEntity(key, fields))`+"\n", g.TranslateType(source.Name))
	result += "}\n"
	result += fmt.Sprintf(`func (c *Client) %ss(where Where) ([]%s, error) {
	type ReturnObj struct {
		Value []%s `+"`"+`json:"value"`+"`"+`
	}

	result := ReturnObj{}

	raw, err := c.getEntities("%s", where)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return nil, err
	}

	return result.Value, nil
}`, g.TranslateType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name), source.Name)
	return result
}

func (g *Generator) GenUpdate(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) Update%s(key Primary%s, entity %s) (*%s, error) {\n", g.TranslateType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += fmt.Sprintf(`	return New%s(c.updateEntity(key, entity))`+"\n", g.TranslateType(source.Name))
	return result + "}"
}

func (g *Generator) GenDelete(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) Delete%s(key Primary%s) error {\n", g.TranslateType(source.Name), g.TranslateType(source.Name))
	result += fmt.Sprintf(`	return c.removeEntity(key)` + "\n")
	return result + "}"
}
