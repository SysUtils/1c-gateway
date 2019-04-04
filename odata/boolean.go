package odata

import (
	"encoding/json"
	"fmt"
)

type Int int

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Int) ImplementsGraphQLType(name string) bool {
	return name == "Int"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *Int) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		tmp := Int(input)
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
func (t Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}
