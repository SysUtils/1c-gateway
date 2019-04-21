package odata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Double float64

func (Double) ImplementsGraphQLType(name string) bool {
	return name == "Double"
}

func (t *Double) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case float32:
		*t = Double(input)
	case float64:
		*t = Double(input)
	case int:
		*t = Double(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Double(val)
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Double) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(t))
}

func (t Double) AsParameter() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 64)
}
