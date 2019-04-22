package odata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Float float64

func (Float) ImplementsGraphQLType(name string) bool {
	return name == "Float"
}

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
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(t))
}

func (t *Float) UnmarshalJSON(b []byte) error {
	val, err := strconv.ParseFloat(strings.Trim(string(b), `"`), 64)
	*t = Float(val)
	return err
}

func (t Float) AsParameter() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 64)
}
