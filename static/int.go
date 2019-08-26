package odata

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"
	"strconv"
	"strings"
)

// Type for Edm.Int
type Int int

// Maps Int to the graphql scalar type in the schema.
func (Int) ImplementsGraphQLType(name string) bool {
	return name == "Int"
}

// A custom unmarshaler for Int type
func (t *Int) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		*t = Int(input)
	case int32:
		*t = Int(input)
	case int64:
		*t = Int(input)
	case float64:
		*t = Int(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Int(val)
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
	return nil
}

// A custom json/graphql marshaller for Int type
func (t Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}

// A custom json unmarshaller for Int type
func (t *Int) UnmarshalJSON(b []byte) error {
	s := string(b)
	val, err := strconv.Atoi(strings.Trim(s, `"`))
	*t = Int(val)
	return err
}

// A custom marshaller to uri query format for Int type
func (t Int) AsParameter() string {
	return strconv.Itoa(int(t))
}
