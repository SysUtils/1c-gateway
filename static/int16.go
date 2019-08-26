package odata

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"
	"strconv"
	"strings"
)

// Type for Edm.Int16
type Int16 int16

// Maps Int to the graphql scalar type in the schema.
func (Int16) ImplementsGraphQLType(name string) bool {
	return name == "Int16"
}

// A custom unmarshaler for Int16 type
func (t *Int16) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		*t = Int16(input)
	case int32:
		*t = Int16(input)
	case int64:
		*t = Int16(input)
	case float64:
		*t = Int16(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Int16(val)
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
	return nil
}

// A custom json/graphql marshaller for Int16 type
func (t Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}

// A custom json unmarshaller for Int16 type
func (t *Int16) UnmarshalJSON(b []byte) error {
	s := string(b)
	val, err := strconv.Atoi(strings.Trim(s, `"`))
	*t = Int16(val)
	return err
}

// A custom marshaller to uri query format for Int16 type
func (t Int16) AsParameter() string {
	return strconv.Itoa(int(t))
}
