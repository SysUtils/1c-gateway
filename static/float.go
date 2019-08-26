package odata

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"
	"strconv"
	"strings"
)

// Type for Edm.Float
type Float float64

// Maps Float to the graphql scalar type in the schema.
func (Float) ImplementsGraphQLType(name string) bool {
	return name == "Float"
}

// A custom graphql unmarshaler for Float type
func (t *Float) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case float32:
		*t = Float(input)
	case float64:
		*t = Float(input)
	case int:
		*t = Float(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Float(val)
	default:
		return errors.Errorf(convertErrorFormat, reflect.TypeOf(input), reflect.TypeOf(*t))
	}
	return nil
}

// A custom json/graphql marshaller for Float type
func (t Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(t))
}

// A custom json unmarshaller for Float type
func (t *Float) UnmarshalJSON(b []byte) error {
	val, err := strconv.ParseFloat(strings.Trim(string(b), `"`), 64)
	*t = Float(val)
	return err
}

// A custom marshaller to uri query format for Float type
func (t Float) AsParameter() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 64)
}
