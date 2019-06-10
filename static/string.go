package odata

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Type for Edm.String
type String string

// Maps String to the graphql scalar type in the schema.
func (String) ImplementsGraphQLType(name string) bool {
	return name == "String"
}

// A custom unmarshaler for String type
func (t *String) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = String(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// A custom json/graphql marshaller for String type
func (t String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

// A custom json unmarshaller for String type
func (t *String) UnmarshalJSON(b []byte) error {
	s := ""
	err := json.Unmarshal(b, &s)
	*t = String(s)
	return err
}

// Escape quotes
func (t String) Escape() string {
	return strings.Replace(string(t), "'", "''", -1)
}

// A custom marshaller to uri query format for String type
func (t String) AsParameter() string {
	return fmt.Sprintf("'%s'", t.Escape())
}
