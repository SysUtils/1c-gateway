package odata

import (
	"encoding/json"
	"fmt"
)

type Float float32

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Float) ImplementsGraphQLType(name string) bool {
	return name == "Float"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *Float) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case float32:
		tmp := Float(input)
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
func (t Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(t))
}
