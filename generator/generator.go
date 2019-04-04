package generator

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Generator struct {
	url          string
	user         string
	pass         string
	Schema       Schema
	types        map[string]Type
	fields       map[string]string
	entities     map[string][]OneCType
	associations map[string]map[string]string
	schemaTmpl   *template.Template
	resolverTmpl *template.Template
	typesTmpl    *template.Template
	clientTmpl   *template.Template
}

func NewGenerator(url, user, pass string, typemap map[string]string, fieldmap map[string]string) (*Generator, error) {
	associations := make(map[string]map[string]string)
	types := map[string]Type{
		"Edm.String":         {"String", nil},
		"Edm.Boolean":        {"Boolean", nil},
		"Edm.Int64":          {"Int", nil},
		"Edm.Guid":           {"Guid", nil},
		"Edm.Double":         {"Float", nil},
		"Edm.Time":           {"Time", nil},
		"Edm.TimeOfDay":      {"Time", nil},
		"Edm.Duration":       {"Time", nil},
		"Edm.Date":           {"Time", nil},
		"Edm.DateTime":       {"Time", nil},
		"Edm.DateTimeOffset": {"Time", nil},
		"Edm.Int16":          {"Int", nil},
		"Edm.Binary":         {"Binary", nil},
		"Edm.Stream":         {"Binary", nil},
		"Edm.Byte":           {"Int", nil},
		"Edm.SByte":          {"Int", nil},
	}
	for key, val := range typemap {
		types["StandardODATA."+key] = Type{val, nil}
	}

	schemaTmpl, err := template.New("").Funcs(template.FuncMap{"BuildEntityRefkeys": BuildEntityRefkeys, "BuildArgs": BuildArgs, "Transliterate": Translit, "BuildFilter": BuildFilter, "ConvertToGqlType": ConvertToGqlType}).Parse(templateSchema)
	if err != nil {
		return nil, err
	}

	resolverTmpl, err := template.New("").Funcs(template.FuncMap{"BuildEntityRefkeys": BuildEntityRefkeys, "BuildArgs": BuildArgs, "Transliterate": Translit, "BuildFilter": BuildFilter, "ConvertToGqlType": ConvertToGqlType}).Parse(templateResolver)
	if err != nil {
		return nil, err
	}

	typesTmpl, err := template.New("").Funcs(template.FuncMap{"BuildEntityRefkeys": BuildEntityRefkeys, "BuildArgs": BuildArgs, "Transliterate": Translit, "BuildFilter": BuildFilter, "ConvertToGqlType": ConvertToGqlType}).Parse(templateTypes)
	if err != nil {
		return nil, err
	}

	clientTmpl, err := template.New("").Funcs(template.FuncMap{"BuildEntityRefkeys": BuildEntityRefkeys, "BuildArgs": BuildArgs, "Transliterate": Translit, "BuildFilter": BuildFilter, "ConvertToGqlType": ConvertToGqlType}).Parse(templateClient)
	if err != nil {
		return nil, err
	}

	return &Generator{url: url, user: user, pass: pass, types: types, fields: fieldmap, associations: associations, schemaTmpl: schemaTmpl, resolverTmpl: resolverTmpl, typesTmpl: typesTmpl, clientTmpl: clientTmpl}, nil
}

func (g *Generator) Bootstrap() error {
	req, err := http.NewRequest("GET", g.url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(g.user, g.pass)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	metadata := Edmx{}
	err = xml.Unmarshal([]byte(data), &metadata)
	if err != nil {
		return err
	}
	g.Schema = metadata.DataServices.Schema
	return nil
}

func (g *Generator) ExtractAssociations() {
	for _, assoc := range g.Schema.Association {
		g.addAssoc(assoc)
	}
}

func (g *Generator) ExtractTypes() {
	for _, entity := range g.Schema.Entities {
		typeName := entity.Name
		name := fmt.Sprintf("%s.%s", g.Schema.Namespace, entity.Name)
		typeName = strings.Replace(typeName, "_", "", -1)
		g.addType(name, typeName)
	}

	for _, entity := range g.Schema.Complexes {
		name := fmt.Sprintf("%s.%s", g.Schema.Namespace, entity.Name)
		typeName := strings.Replace(entity.Name, "_", "", -1)
		g.addType(name, typeName)
	}
}

func (g *Generator) PrepareComplexes() {
	g.Schema.ComplexImports = make(map[string]bool)
	for complexId, entity := range g.Schema.Complexes {
		val := g.translateType("StandardODATA." + entity.Name)
		g.Schema.Complexes[complexId].CamelName = val.Name
		for propId, prop := range entity.Properties {
			val := g.translateType(prop.Type)
			g.Schema.Complexes[complexId].Properties[propId].CamelType = val.Name
			if val, ok := g.fields[prop.Name]; ok {
				g.Schema.Complexes[complexId].Properties[propId].CamelName = val
			} else {
				g.Schema.Complexes[complexId].Properties[propId].CamelName = strings.Replace(prop.Name, "_", "", -1)
			}
			for _, imp := range val.Imports {
				g.Schema.ComplexImports[imp] = true
			}
		}
	}
}

func (g *Generator) PrepareEntities() {
	g.Schema.EntityImports = make(map[string]bool)
	for entityId, entity := range g.Schema.Entities {
		fullname := "StandardODATA." + entity.Name
		for _, function := range g.Schema.Functions {
			if function.Parameters[0].Type == fullname {
				g.Schema.Entities[entityId].Functions = append(g.Schema.Entities[entityId].Functions, function)
			}
		}

		val := g.translateType(fullname)
		g.Schema.Entities[entityId].CamelName = val.Name
		for propId, prop := range entity.Properties {
			val := g.translateType(prop.Type)
			g.Schema.Entities[entityId].Properties[propId].CamelType = val.Name
			if val, ok := g.fields[prop.Name]; ok {
				g.Schema.Entities[entityId].Properties[propId].CamelName = val
			} else {
				g.Schema.Entities[entityId].Properties[propId].CamelName = strings.Replace(prop.Name, "_", "", -1)
			}
			for _, imp := range val.Imports {
				g.Schema.EntityImports[imp] = true
			}
			for keyId, key := range entity.Keys {
				if key.Name == prop.Name {
					if val, ok := g.fields[prop.Name]; ok {
						g.Schema.Entities[entityId].Keys[keyId].CamelName = val
					} else {
						g.Schema.Entities[entityId].Keys[keyId].CamelName = strings.Replace(key.Name, "_", "", -1)
					}

					g.Schema.Entities[entityId].Keys[keyId].Type = val.Name
				}
			}
		}
		for navId, nav := range entity.Navigations {
			val := g.translateType(g.associations[nav.Type][nav.ToRole])
			g.Schema.Entities[entityId].Navigations[navId].CamelType = val.Name
			g.Schema.Entities[entityId].Navigations[navId].CamelName = strings.Replace(nav.Name, "_", "", -1)
			for _, imp := range val.Imports {
				g.Schema.EntityImports[imp] = true
			}
		}
	}
}

func (g *Generator) PrepareMethods() {
	for i, f := range g.Schema.Functions {
		if f.Type != "" {
			g.Schema.Functions[i].Type = g.translateType(f.Type).Name
		}

		if g.Schema.Functions[i].IsBindable {
			g.Schema.Functions[i].Parameters = g.Schema.Functions[i].Parameters[1:]
		}
		for ii, p := range g.Schema.Functions[i].Parameters {
			if p.Type != "" {
				g.Schema.Functions[i].Parameters[ii].Type = g.translateType(p.Type).Name
			}
		}
	}

	for entityIndex := range g.Schema.Entities {
		for i, f := range g.Schema.Entities[entityIndex].Functions {
			if f.Type != "" {
				g.Schema.Entities[entityIndex].Functions[i].Type = g.translateType(f.Type).Name
			}

			if g.Schema.Entities[entityIndex].Functions[i].IsBindable {
				g.Schema.Entities[entityIndex].Functions[i].Parameters = g.Schema.Entities[entityIndex].Functions[i].Parameters[1:]
			}
		}
	}

	fmt.Print("dsafh")
}

func assetToFile(asset, path string) error {
	data, err := Asset(asset)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func templateToFile(tmpl *template.Template, data interface{}, path string) error {
	err := os.MkdirAll(path[:strings.LastIndex(path, "/")], os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) Write() error {
	err := os.MkdirAll("odata/", os.ModePerm)
	if err != nil {
		return err
	}

	err = templateToFile(g.schemaTmpl, g.Schema, "odata/schema.gql")
	if err != nil {
		return err
	}

	err = templateToFile(g.resolverTmpl, g.Schema, "odata/Resolver.go")
	if err != nil {
		return err
	}

	err = templateToFile(g.typesTmpl, g.Schema, "odata/Types.go")
	if err != nil {
		return err
	}

	err = templateToFile(g.clientTmpl, g.Schema, "odata/Entities.go")
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) addType(key, val string) {
	if _, ok := g.types[key]; !ok {
		g.types[key] = Type{val, nil}
	}
}

func (g *Generator) translateType(t string) Type {
	if val, ok := g.types[t]; ok {
		return val
	}
	if strings.HasPrefix(t, "Collection(") && strings.HasSuffix(t, ")") {
		res := g.translateType(t[11 : len(t)-1])
		return Type{"[]" + res.Name, res.Imports}
	}
	log.Fatalf("Type not found: %s\n", t)
	return Type{t, nil}
}

func (g *Generator) addAssoc(assoc Association) {
	name := fmt.Sprintf("%s.%s", g.Schema.Namespace, assoc.Name)
	if g.associations[name] == nil {
		g.associations[name] = make(map[string]string)
	}
	for _, end := range assoc.Ends {
		g.associations[name][end.Role] = end.Type
	}
}
