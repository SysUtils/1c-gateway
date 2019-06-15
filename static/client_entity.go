package odata

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (c *Client) getEntity(keys IPrimaryKey, fields []string) ([]byte, error) {
	uri := "/" + url.PathEscape(keys.APIEntityType())
	uri += fmt.Sprintf("(%s)", keys.Serialize()) // Unique key
	uri += "?$format=json"
	if len(fields) > 0 {
		uri += fmt.Sprintf("&$select=%s", url.PathEscape(strings.Join(fields, ","))) // Fields
	}

	return c.get(uri)
}

func (c *Client) getEntities(name string, where Where) ([]byte, error) {
	uri := "/" + url.PathEscape(name)
	uri += "?$format=json&"
	uri += where.Serialize()

	return c.get(uri)
}

// Returns json representation of entity's NavigationProperty
func (c *Client) GetEntityNavigaion(key IPrimaryKey, property string) ([]byte, error) {
	uri := "/" + url.PathEscape(key.APIEntityType())
	uri += fmt.Sprintf("(%s)", url.PathEscape(key.Serialize())) // Unique key
	uri += "/" + url.PathEscape(property)
	uri += "?$format=json&"

	body, err := c.get(uri)
	if err != nil {
		if err.Error() == "404 Not found\nBody:\n" {
			return nil, nil
		}
	}

	return body, err
}

// Execute entity's method and return its output in json
func (c *Client) ExecuteEntityMethod(key IPrimaryKey, function IFunction) ([]byte, error) {
	uri := "/" + url.PathEscape(key.APIEntityType())
	uri += fmt.Sprintf("(%s)", url.PathEscape(key.Serialize())) // Unique key
	uri += "/" + function.Name()
	uri += "?$format=json&"
	uri += function.Parameters()

	return c.post(uri, nil)
}

func (c *Client) updateEntity(key IPrimaryKey, entity IEntity) ([]byte, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	uri := "/" + url.PathEscape(entity.APIEntityType())

	uri += fmt.Sprintf("(%s)", url.PathEscape(key.Serialize())) // Unique key
	uri += "?$format=json"

	return c.patch(uri, data)
}

func (c *Client) removeEntity(key IPrimaryKey) error {
	uri := "/" + url.PathEscape(key.APIEntityType())

	uri += fmt.Sprintf("(%s)", url.PathEscape(key.Serialize())) // Unique key
	uri += "?$format=json"

	return c.delete(uri)
}

func (c *Client) createEntity(entity IEntity) ([]byte, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	uri := "/" + url.PathEscape(entity.APIEntityType()) + "?$format=json"

	return c.post(uri, data)
}
