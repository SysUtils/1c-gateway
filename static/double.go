package odata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Type for Edm.Double
type Double float64

// Maps Double to the graphql scalar type in the schema.
func (Double) ImplementsGraphQLType(name string) bool {
	return name == "Double"
}

// A custom graphql unmarshaler for Double type
func (t *Double) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case float32:
		*t = Double(input)
	case float64:
		*t = Double(input)
	case int:
		*t = Double(input)
	case string:
		val, err := strconv.ParseFloat(strings.Trim(input, `"`), 64)
		if err != nil {
			return err
		}
		*t = Double(val)
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

// A custom json/graphql marshaller for Double type
func (t Double) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(t))
}

// A custom json unmarshaller for Double type
func (t *Double) UnmarshalJSON(b []byte) error {
	val, err := strconv.ParseFloat(strings.Trim(string(b), `"`), 64)
	*t = Double(val)
	return err
}

// A custom marshaller to uri query format for Double type
func (t Double) AsParameter() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 64)
}
