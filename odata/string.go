package odata

import (
	"encoding/json"
	"fmt"
)

type String string

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (String) ImplementsGraphQLType(name string) bool {
	return name == "String"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *String) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		tmp := String(input)
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
func (t String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}
