package native

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genTypes(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.genType(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) genType(source shared.OneCType) string {
	result := g.genTypeStruct(source)
	result += "\n"
	result += g.genTypeFunc(source)
	result += "\n"
	result += g.genPrimaryKeyFunc(source)
	result += "\n"
	result += g.genNew(source)
	result += "\n"
	result += g.genCreate(source)
	result += "\n"
	result += g.genRead(source)
	result += "\n"
	result += g.genUpdate(source)
	result += "\n"
	result += g.genDelete(source)
	return result
}

func (g *Generator) genTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type %s struct {\n", g.translateType(source.Name))
	result += "	Client *Client\n"
	for _, prop := range source.Properties {
		result += "	"
		result += g.translateName(prop.Name)
		result += " *"
		/*if prop.Nullable {
			result += ""
		}*/
		result += g.translateType(prop.Type)
		result += " `"
		result += fmt.Sprintf(`json:"%s,omitempty"`, prop.Name)
		result += "`"
		result += "\n"
	}
	return result + "}"
}

func (g *Generator) genNew(source shared.OneCType) string {
	result := fmt.Sprintf("func New%s(data []byte, prevError error, client *Client) (*%s, error) {\n", g.translateType(source.Name), g.translateType(source.Name))
	result += fmt.Sprintf(`	if prevError != nil {
		return nil, prevError
	}
	result := new(%s)
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	result.Client = client
	return result, nil
}`, g.translateType(source.Name))
	return result
}

func (g *Generator) genTypeFunc(source shared.OneCType) string {
	result := fmt.Sprintf("func (%s) APIEntityType() string {\n", g.translateType(source.Name))
	result += fmt.Sprintf(`	return "%s"
}`, source.Name)
	return result
}

func (g *Generator) genNavigationProperties(source shared.OneCType) string {
	result := fmt.Sprintf("func (%s) () string {\n", g.translateType(source.Name))
	result += fmt.Sprintf(`	return "%s"
}`, source.Name)
	return result
}

func (g *Generator) genPrimaryKeyFunc(source shared.OneCType) string {
	result := fmt.Sprintf("func (e %s) PrimaryKey () (*Primary%s, error) {\n", g.translateType(source.Name), g.translateType(source.Name))
	for _, key := range source.Keys {
		result += fmt.Sprintf("	if e.%s == nil {", g.translateName(key.Name))
		result += fmt.Sprintf("		return nil,  errors.New(`e.%s must be not null`)", g.translateName(key.Name))
		result += fmt.Sprintf("	}")
	}
	result += fmt.Sprintf("	return Primary%s {", g.translateType(source.Name))
	for _, key := range source.Keys {
		result += fmt.Sprintf("%s: *(e.%s),", g.translateName(key.Name), g.translateName(key.Name))
	}
	result += fmt.Sprintf(`}
}`)
	return result
}

func (g *Generator) genCreate(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) Create%s(entity %s) (*%s, error) {\n", g.translateType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	result += fmt.Sprintf(`	src, err := c.createEntity(entity)` + "\n")
	result += fmt.Sprintf(`	return New%s(src, err, c)`+"\n", g.translateType(source.Name))
	return result + "}"
}

func (g *Generator) genRead(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) %s(key Primary%s, fields []string) (*%s, error) {\n", g.translateType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	result += fmt.Sprintf(`	src, err := c.getEntity(key, fields)` + "\n")
	result += fmt.Sprintf(`	return New%s(src, err, c)`+"\n", g.translateType(source.Name))
	result += "}\n"
	result += fmt.Sprintf(`func (c *Client) %ss(where Where) (*[]%s, error) {
	type ReturnObj struct {
		Value []%s `+"`"+`json:"value"`+"`"+`
	}

	result := ReturnObj{}

	raw, err := c.getEntities("%s", where)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &result)
	if err != nil {
		return nil, err
	}
	for i := range result.Value {
		result.Value[i].Client = c;
	}

	return &result.Value, nil
}`, g.translateType(source.Name), g.translateType(source.Name), g.translateType(source.Name), source.Name)
	return result
}

func (g *Generator) genUpdate(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) Update%s(key Primary%s, entity %s) (*%s, error) {\n", g.translateType(source.Name), g.translateType(source.Name), g.translateType(source.Name), g.translateType(source.Name))
	result += fmt.Sprintf(`	src, err := c.updateEntity(key, entity)` + "\n")
	result += fmt.Sprintf(`	return New%s(src, err, c)`+"\n", g.translateType(source.Name))
	return result + "}"
}

func (g *Generator) genDelete(source shared.OneCType) string {
	result := fmt.Sprintf("func (c *Client) Delete%s(key Primary%s) error {\n", g.translateType(source.Name), g.translateType(source.Name))
	result += fmt.Sprintf(`	return c.removeEntity(key)` + "\n")
	return result + "}"
}
