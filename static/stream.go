package odata

import (
	"github.com/go-errors/errors"
	"reflect"
)

// Type for Edm.Stream
type Stream string

// Maps this Go type to the graphql scalar type in the schema.
func (Stream) ImplementsGraphQLType(name string) bool {
	return name == "Stream"
}

// A custom graphql unmarshaller for Binary type
func (t *Stream) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = Stream(input)
		return nil
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
}

// A custom marshaller to uri query format
func (t Stream) AsParameter() string {
	return `'` + string(t) + `'`
}
