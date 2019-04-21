package odata

import (
	"encoding/json"
	"fmt"
)

type Guid string

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Guid) ImplementsGraphQLType(name string) bool {
	return name == "Guid"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *Guid) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = Guid(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// MarshalJSON is a custom marshaler for Time
//
// This function will be called whenever you
// query for fields that use the Time type
func (t Guid) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t Guid) AsParameter() string {
	return `guid'` + string(t) + `'`
}
