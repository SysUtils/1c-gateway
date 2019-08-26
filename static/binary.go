package odata

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"
	"strings"
)

// Type for Edm.Binary
type Binary []byte

// Maps this Go type to the graphql scalar type in the schema.
func (Binary) ImplementsGraphQLType(name string) bool {
	return name == "Binary"
}

// A custom graphql unmarshaller for Binary type
func (t *Binary) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		res, err := b64.StdEncoding.DecodeString(input)
		*t = res
		return err
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
}

// A custom json/graphql marshaller for Binary type
func (t Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(b64.StdEncoding.EncodeToString(t))
}

// A custom json unmarshaller for Binary type
func (t *Binary) UnmarshalJSON(b []byte) error {
	val, err := b64.StdEncoding.DecodeString(strings.Trim(string(b), `"`))
	*t = val
	return err
}

// A custom marshaller to uri query format
func (t Binary) AsParameter() string {
	return `'` + b64.StdEncoding.EncodeToString(t) + `'`
}
