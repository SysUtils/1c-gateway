package odata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Int int

func (Int) ImplementsGraphQLType(name string) bool {
	return name == "Int"
}

func (t *Int) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		*t = Int(input)
	case int32:
		*t = Int(input)
	case int64:
		*t = Int(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Int(val)
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}

func (t *Int) UnmarshalJSON(b []byte) error {
	s := string(b)
	val, err := strconv.Atoi(strings.Trim(s, `"`))
	*t = Int(val)
	return err
}

func (t Int) AsParameter() string {
	return strconv.Itoa(int(t))
}
