package odata

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"
)

// Type for Edm.Boolean
type Boolean bool

// Maps Boolean to the graphql scalar type in the schema.
func (Boolean) ImplementsGraphQLType(name string) bool {
	return name == "Boolean"
}

// A custom graphql unmarshaller for Boolean type
func (t *Boolean) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case bool:
		*t = Boolean(input)
		return nil
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
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
