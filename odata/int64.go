package odata

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Int64 int64

func (Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

func (t *Int64) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case int:
		*t = Int64(input)
	case int32:
		*t = Int64(input)
	case int64:
		*t = Int64(input)
	case string:
		val, err := strconv.Atoi(strings.Trim(input, `"`))
		if err != nil {
			return err
		}
		*t = Int64(val)
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}

func (t *Int64) UnmarshalJSON(b []byte) error {
	s := string(b)
	val, err := strconv.ParseInt(strings.Trim(s, `"`), 10, 64)
	*t = Int64(val)
	return err
}

func (t Int64) AsParameter() string {
	return strconv.FormatInt(int64(t), 10)
}
