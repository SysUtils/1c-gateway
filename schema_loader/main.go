// Package schema_loader provides downloader for metadata
package schema_loader

import (
	"encoding/xml"
	"github.com/SysUtils/1c-gateway/shared"
	"io/ioutil"
	"net/http"
)

type SchemaLoader struct {
	url  string
	user string
	pass string
}

// Returns initialized schema loader
func NewSchemaLoader(url, user, pass string) *SchemaLoader {
	return &SchemaLoader{url: url, user: user, pass: pass}
}

// Download metadata and extract schema object from it
func (g *SchemaLoader) Load() (*shared.Schema, error) {
	req, err := http.NewRequest("GET", g.url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(g.user, g.pass)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	metadata := shared.Edmx{}
	err = xml.Unmarshal([]byte(data), &metadata)
	if err != nil {
		return nil, err
	}
	return &metadata.DataServices.Schema, nil
}
