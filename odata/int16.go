package odata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Int16 int16

func (Int16) ImplementsGraphQLType(name string) bool {
	return name == "Int16"
}

func (t *Int16) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		*t = Int16(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Int16(val)
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}

func (t Int16) AsParameter() string {
	return strconv.Itoa(int(t))
}
