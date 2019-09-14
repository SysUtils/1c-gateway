package odata

import (
	"github.com/go-errors/errors"
	"reflect"
)

// Type for Edm.Binary
type Binary string

// Maps this Go type to the graphql scalar type in the schema.
func (Binary) ImplementsGraphQLType(name string) bool {
	return name == "Binary"
}

// A custom graphql unmarshaller for Binary type
func (t *Binary) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = Binary(input)
		return nil
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
}

// A custom marshaller to uri query format
func (t Binary) AsParameter() string {
	return `'` + string(t) + `'`
}
