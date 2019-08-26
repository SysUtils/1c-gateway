package odata

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"
	"strings"
)

// Type for Edm.Stream
type Stream []byte

// Maps Stream to the graphql scalar type in the schema.
func (Stream) ImplementsGraphQLType(name string) bool {
	return name == "Stream"
}

// A custom unmarshaler for Stream type
func (t *Stream) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		res, _ := b64.StdEncoding.DecodeString(input)
		*t = Stream(res)
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
	return nil
}

// A custom json/graphql marshaller for Stream type
func (t Stream) MarshalJSON() ([]byte, error) {
	return json.Marshal(b64.StdEncoding.EncodeToString(t))
}

// A custom json unmarshaller for Stream type
func (t *Stream) UnmarshalJSON(b []byte) error {
	val, err := b64.StdEncoding.DecodeString(strings.Trim(string(b), `"`))
	*t = Stream(val)
	return err
}

// A custom marshaller to uri query format for Stream type
func (t Stream) AsParameter() string {
	return `'` + b64.StdEncoding.EncodeToString(t) + `'`
}
