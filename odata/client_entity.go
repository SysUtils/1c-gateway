package odata

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (c *Client) getEntity(name string, parameters []string, fields []string) (string, error) {
	uri := "/" + url.PathEscape(name)
	if len(parameters) > 0 {
		uri += fmt.Sprintf("(%s)", url.PathEscape(strings.Join(parameters, ","))) // Unique key
	} else {
		return "", ErrEmptyParams
	}
	uri += "?$format=json"
	if len(fields) > 0 {
		uri += fmt.Sprintf("&$select=%s", url.PathEscape(strings.Join(fields, ","))) // Fields
	}

	return c.get(uri)
}

func (c *Client) getEntities(name string, where Where) (string, error) {
	uri := "/" + url.PathEscape(name)
	uri += "?$format=json&"

	uri += SerializeWhere(where)

	return c.get(uri)
}

func (c *Client) executeEntityMethod(entityName string, keys []string, name string, parameters map[string]string) (string, error) {
	uri := "/" + url.PathEscape(name)

	if len(parameters) > 0 {
		uri += fmt.Sprintf("(%s)", url.PathEscape(strings.Join(keys, ","))) // Unique key
	} else {
		return "", ErrEmptyParams
	}
	uri += "?$format=json"
	for key, val := range parameters {
		uri += fmt.Sprintf("&%s=%s", url.QueryEscape(key), url.QueryEscape(val))
	}

	return c.post(uri, nil)
}

func (c *Client) updateEntity(name string, keys []string, entity interface{}) (string, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return "", err
	}

	uri := "/" + url.PathEscape(name)

	if len(keys) > 0 {
		uri += fmt.Sprintf("(%s)", url.PathEscape(strings.Join(keys, ","))) // Unique key
	} else {
		return "", ErrEmptyParams
	}
	uri += "?$format=json"

	return c.patch(uri, data)
}

func (c *Client) removeEntity(name string, keys []string) error {
	uri := "/" + url.PathEscape(name)

	if len(keys) > 0 {
		uri += fmt.Sprintf("(%s)", url.PathEscape(strings.Join(keys, ","))) // Unique key
	} else {
		return ErrEmptyParams
	}
	uri += "?$format=json"

	return c.delete(uri)
}

func (c *Client) createEntity(name string, entity interface{}) (string, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return "", err
	}
	uri := "/" + url.PathEscape(name) + "?$format=json"

	return c.post(uri, data)
}
