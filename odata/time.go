package odata

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Time time.Time

func (Time) ImplementsGraphQLType(name string) bool {
	return name == "Time"
}

func (t *Time) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		*t = Time(input)
		return nil
	case string:
		val, err := time.Parse(time.RFC3339, input)
		*t = Time(val)
		return err
	case int:
		val := time.Unix(int64(input), 0)
		*t = Time(val)
		return nil
	case float64:
		val := time.Unix(int64(input), 0)
		*t = Time(val)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t Time) MarshalJSON() ([]byte, error) {
	val := time.Time(t).Format(time.RFC3339)
	return json.Marshal(val)
}

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	val, err := time.Parse(time.RFC3339, s)
	*t = Time(val)
	return err
}

func (t Time) AsParameter() string {
	val := time.Time(t).Format(time.RFC3339)
	return fmt.Sprintf("'%s'", val)
}
