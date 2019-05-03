package odata

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

const timeformat = "2006-01-02T15:04:05"

func (DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

func (t *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		*t = DateTime(input)
		return nil
	case string:
		val, err := time.Parse(timeformat, input)
		*t = DateTime(val)
		return err
	case int:
		val := time.Unix(int64(input), 0)
		*t = DateTime(val)
		return nil
	case float64:
		val := time.Unix(int64(input), 0)
		*t = DateTime(val)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	val := time.Time(t).Format(timeformat)
	return json.Marshal(val)
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	val, err := time.Parse(timeformat, strings.Trim(string(b), `"`))
	*t = DateTime(val)
	return err
}

func (t DateTime) AsParameter() string {
	val := time.Time(t).Format(timeformat)
	return fmt.Sprintf("'%s'", val)
}
