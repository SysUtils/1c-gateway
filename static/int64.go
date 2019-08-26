package odata

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"
	"strconv"
	"strings"
)

// Type for Edm.Int64
type Int64 int64

// Maps Int64 to the graphql scalar type in the schema.
func (Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

// A custom unmarshaler for Int64 type
func (t *Int64) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		*t = Int64(input)
	case int32:
		*t = Int64(input)
	case int64:
		*t = Int64(input)
	case float64:
		*t = Int64(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Int64(val)
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
	return nil
}

// A custom json/graphql marshaller for Int64 type
func (t Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}

// A custom json unmarshaller for Int64 type
func (t *Int64) UnmarshalJSON(b []byte) error {
	s := string(b)
	val, err := strconv.ParseInt(strings.Trim(s, `"`), 10, 64)
	*t = Int64(val)
	return err
}

// A custom marshaller to uri query format for Int64 type
func (t Int64) AsParameter() string {
	return strconv.FormatInt(int64(t), 10)
}
