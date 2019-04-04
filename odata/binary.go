package odata

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

type Binary []byte

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Binary) ImplementsGraphQLType(name string) bool {
	return name == "Binary"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *Binary) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		res, _ := b64.StdEncoding.DecodeString(input)
		tmp := Binary(res)
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
func (t Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}
