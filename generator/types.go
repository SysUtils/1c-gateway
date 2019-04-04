package generator

import (
	"encoding/xml"
)

type Type struct {
	Name    string
	Imports []string
}

type PropertyRef struct {
	Name       string `xml:"Name,attr"`
	CamelName  string
	Type       string
	FilterType string
}

type Association struct {
	Name string `xml:"Name,attr"`
	Ends []End  `xml:"End"`
}

type End struct {
	Role string `xml:"Role,attr"`
	Type string `xml:"Type,attr"`
}

type Property struct {
	Name      string `xml:"Name,attr"`
	CamelName string
	Type      string `xml:"Type,attr"`
	CamelType string
	Nullable  bool `xml:"Nullable,attr"`
}

type Parameter struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type Function struct {
	Name       string      `xml:"Name,attr"`
	Type       string      `xml:"ReturnType,attr"`
	IsBindable bool        `xml:"IsBindable,attr"`
	Parameters []Parameter `xml:"Parameter"`
}

type NavProp struct {
	Name      string `xml:"Name,attr"`
	CamelName string
	Type      string `xml:"Relationship,attr"`
	CamelType string
	FromRole  string `xml:"FromRole,attr"`
	ToRole    string `xml:"ToRole,attr"`
}

type OneCType struct {
	Name        string `xml:"Name,attr"`
	CamelName   string
	Keys        []PropertyRef `xml:"Key>PropertyRef"`
	Properties  []Property    `xml:"Property"`
	Navigations []NavProp     `xml:"NavigationProperty"`
	Functions   []Function
	WrongKeys   bool
}

type Schema struct {
	Name           xml.Name
	Namespace      string     `xml:"Namespace,attr"`
	Entities       []OneCType `xml:"EntityType"`
	Complexes      []OneCType `xml:"ComplexType"`
	Association    []Association
	Functions      []Function `xml:"EntityContainer>FunctionImport"`
	ComplexImports map[string]bool
	EntityImports  map[string]bool
}

type NodeValue struct {
	Type       int
	Name       string
	Properties map[string]Property
}

type DataServices struct {
	Schema Schema
}

type Edmx struct {
	DataServices DataServices
}
