package odata

import (
	"encoding/json"
	"fmt"
)

// Type for Edm.Boolean
type Boolean bool

// Maps Boolean to the graphql scalar type in the schema.
func (Boolean) ImplementsGraphQLType(name string) bool {
	return name == "Boolean"
}

// A custom graphql unmarshaler for Boolean type
func (t *Boolean) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case bool:
		*t = Boolean(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// A custom json/graphql marshaller for Boolean type
func (t Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(t))
}

// A custom json unmarshaller for Boolean type
func (t *Boolean) UnmarshalJSON(b []byte) error {
	*t = Boolean(string(b) == "true")
	return nil
}

// A custom marshaller to uri query format for Boolean type
func (t Boolean) AsParameter() string {
	if t {
		return "true"
	} else {
		return "false"
	}
}
