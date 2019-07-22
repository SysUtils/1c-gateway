// Package shared provides shared types
package shared

type PropertyRef struct {
	Name string `xml:"Name,attr"`
	Type string
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
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Type,attr"`
	Nullable bool   `xml:"Nullable,attr"`
}

type Parameter struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

type Function struct {
	Name            string      `xml:"Name,attr"`
	Type            string      `xml:"ReturnType,attr"`
	IsBindable      bool        `xml:"IsBindable,attr"`
	IsSideEffecting bool        `xml:"IsSideEffecting,attr"`
	Parameters      []Parameter `xml:"Parameter"`
}

type NavProp struct {
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Relationship,attr"`
	FromRole string `xml:"FromRole,attr"`
	ToRole   string `xml:"ToRole,attr"`
}

type OneCType struct {
	Name        string        `xml:"Name,attr"`
	Keys        []PropertyRef `xml:"Key>PropertyRef"`
	Properties  []Property    `xml:"Property"`
	Navigations []NavProp     `xml:"NavigationProperty"`
}

type Schema struct {
	Name        string     `xml:"Name,attr"`
	Namespace   string     `xml:"Namespace,attr"`
	Entities    []OneCType `xml:"EntityType"`
	Complexes   []OneCType `xml:"ComplexType"`
	Association []Association
	Functions   []Function `xml:"EntityContainer>FunctionImport"`
}

type DataServices struct {
	Schema Schema
}

type Edmx struct {
	DataServices DataServices
}
