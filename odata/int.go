package odata

import (
	"encoding/json"
	"fmt"
)

type Boolean bool

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Boolean) ImplementsGraphQLType(name string) bool {
	return name == "Boolean"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *Boolean) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case bool:
		tmp := Boolean(input)
		t = &tmp
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// MarshalJSON is a custom marshaler for Time
//
// This function will be called whenever you
// query for fields that use the Time type
func (t Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(t))
}
