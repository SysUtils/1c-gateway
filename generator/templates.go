package generator

//go:generate go-bindata -pkg $GOPACKAGE static/

var templateFields = `
package Fields{{.CamelName}}
{{ range .Properties }}
	const {{.CamelName}} = "{{ .Name }}"
{{- end }}
`

var templateSchema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
{{- range .Entities }}
{{- $name := .Name }}
{{- $camelname := Transliterate .CamelName }}
	{{ $camelname }}: {{ $camelname }}
	{{ $camelname }}s: [{{ $camelname }}!]
{{- end }}
}

type Mutation {
{{- range .Entities }}
{{- $name := .Name }}
{{- $camelname := Transliterate .CamelName }}
	Update{{ $camelname }}: {{ $camelname }}
	Remove{{ $camelname }}: {{ $camelname }}
{{- end }}
}

{{ range .Complexes }}
{{ $name := .Name -}}
{{ $camelname := Transliterate .CamelName -}}
type {{ $camelname }} {
{{- range .Properties -}}
{{- if .Nullable }}
	{{Transliterate .CamelName}}: {{Transliterate (ConvertToGqlType .CamelType) }}
{{- else }}
	{{Transliterate .CamelName}}: {{Transliterate (ConvertToGqlType .CamelType) }}!
{{- end }}
{{- end }}
}
{{ end }}

{{ range .Entities }}
{{ $name := .Name -}}
{{ $camelname := Transliterate .CamelName -}}
type {{ $camelname }} {
{{- range .Properties -}}
{{- if .Nullable }}
	{{ Transliterate .CamelName}}: {{ Transliterate (ConvertToGqlType .CamelType) }}
{{- else }}
	{{ Transliterate .CamelName}}: {{ Transliterate (ConvertToGqlType .CamelType) }}!
{{- end }}
{{- end }}
}
{{ end }}
`

var templateTypes = `
package odata

import (
	"encoding/json"
	"fmt"
	"net/url"
)

{{ range .Complexes }}
{{ $name := .Name -}}
{{ $camelname := Transliterate .CamelName -}}
type {{ $camelname }} struct {
{{- range .Properties -}}
{{- if .Nullable }}
	{{Transliterate .CamelName}} *{{Transliterate .CamelType }} ` + "`" + `xml:"{{.Name}}" json:"{{.Name}},omitempty"` + "`" + `
{{- else }}
	{{Transliterate .CamelName}} {{Transliterate .CamelType }} ` + "`" + `xml:"{{.Name}}" json:"{{.Name}},omitempty"` + "`" + `
{{- end }}
{{- end }}
}
{{ end }}

{{ range .Entities }}
{{ $entity := . -}}
{{ $name := .Name -}}
{{ $camelname := Transliterate .CamelName -}}
type {{ $camelname }} struct {
{{- range .Properties -}}
{{- if .Nullable }}
	{{ Transliterate .CamelName}} *{{ Transliterate .CamelType }} ` + "`" + `xml:"{{.Name}}" json:"{{.Name}},omitempty"` + "`" + `
{{- else }}
	{{ Transliterate .CamelName}} {{ Transliterate .CamelType }} ` + "`" + `xml:"{{.Name}}" json:"{{.Name}},omitempty"` + "`" + `
{{- end }}
{{- end }}
}
{{ range .Functions }}
{{ if eq .Type "" -}}
func (entity {{ $camelname }}) {{.Name}}({{- range .Parameters -}}
{{ Transliterate .Name}} {{ Transliterate .Type}},
{{- end -}}) error {
	return nil
}
{{- else -}}
func (entity {{ $camelname }}) {{.Name}}(client Client, {{- range .Parameters -}}
{{ Transliterate .Name}} {{ Transliterate .Type}},
{{- end -}}) ({{ Transliterate .Type}}, error) {
	var result {{ Transliterate .Type}}
	args := {{ BuildArgs .Parameters }}
	keys := {{ BuildEntityRefkeys $entity.Keys }}
	raw, err := client.executeEntityMethod("{{$entity.Name}}", keys, "{{.Name}}", args)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
{{ end }}
{{ end }}
{{ end }}
`

var templateResolver = `
package odata

type Resolver struct {}

{{- range .Entities }}
{{- $name := .Name }}
{{- $camelname := Transliterate .CamelName }}
func (_ *Resolver) Update{{ $camelname }}() (*{{ $camelname }}, error) {
	return nil, nil
}
func (_ *Resolver) Remove{{ $camelname }}() (*{{ $camelname }}, error) {
	return nil, nil
}

func (_ *Resolver) {{ $camelname }}() (*{{ $camelname }}, error) {
	return nil, nil
}
func (_ *Resolver) {{ $camelname }}s() (*[]{{ $camelname }}, error) {
	return nil, nil
}
{{end}}`

var templateClient = `
package odata

import (
	"encoding/json"
	"fmt"
	"net/url"
)

{{- range .Entities }}
{{- $name := .Name }}
{{- $camelname := Transliterate .CamelName }}

{{ if not .WrongKeys }}
func (c *Client) {{ $camelname }}({{- range .Keys -}}
{{.CamelName}} {{.Type}},
{{- end -}}) (*{{ $camelname }}, error) {
	args := {{ BuildFilter .Keys }}
	raw, err := c.getEntity("{{.Name}}", args, []string {})
	if err != nil {
		return nil, err
	}
	result := {{ $camelname }}{}
	err = json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Update{{ $camelname }}({{- range .Keys -}}
{{.CamelName}} {{.Type}},
{{- end -}} entity interface{}) (*{{ $camelname }}, error) {
	args := {{ BuildFilter .Keys }}
	raw, err := c.updateEntity("{{.Name}}", args, entity)
	if err != nil {
		return nil, err
	}
	result := {{ $camelname }} {}
	err = json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
func (c *Client) Remove{{ $camelname }}({{- range .Keys -}}
{{.CamelName}} {{.Type}},
{{- end -}}) error {
	args := {{ BuildFilter .Keys }}
	return c.removeEntity("{{.Name}}", args)
}
{{end}}

func (c *Client) Create{{ $camelname }} (data {{ $camelname }}) (*{{ $camelname }}, error) {
	raw, err := c.createEntity("{{.Name}}", data)
	if err != nil {
		return nil, err
	}
	result := {{ $camelname }} {}
	err = json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) {{ $camelname }}s(where Where) ([]{{ $camelname }}, error) {
	type ReturnObj struct {
		Value []{{ $camelname }} ` + "`" + `json:"value"` + "`" + `
	}

	result := ReturnObj{}

	raw, err := c.getEntities("{{.Name}}", where)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return nil, err
	}

	return result.Value, nil
}
{{end}}`
