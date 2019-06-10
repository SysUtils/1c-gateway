package odata

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Type for Edm.Binary
type Binary []byte

// Maps this Go type to the graphql scalar type in the schema.
func (Binary) ImplementsGraphQLType(name string) bool {
	return name == "Binary"
}

// A custom graphql unmarshaler for Binary type
func (t *Binary) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		res, err := b64.StdEncoding.DecodeString(input)
		*t = Binary(res)
		return err
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

// A custom json/graphql marshaller for Binary type
func (t Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(b64.StdEncoding.EncodeToString(t))
}

// A custom json unmarshaller for Binary type
func (t *Binary) UnmarshalJSON(b []byte) error {
	val, err := b64.StdEncoding.DecodeString(strings.Trim(string(b), `"`))
	*t = Binary(val)
	return err
}

// A custom marshaller to uri query format
func (t Binary) AsParameter() string {
	return `'` + b64.StdEncoding.EncodeToString(t) + `'`
}
