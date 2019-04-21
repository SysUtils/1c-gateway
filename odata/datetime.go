package odata

import (
	"encoding/json"
	"fmt"
	"time"
)

type DateTime time.Time

func (DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

func (t *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		*t = DateTime(input)
		return nil
	case string:
		val, err := time.Parse(time.RFC3339, input)
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
	val := time.Time(t).Format(time.RFC3339)
	return json.Marshal(val)
}

func (t DateTime) AsParameter() string {
	val := time.Time(t).Format(time.RFC3339)
	return fmt.Sprintf("'%s'", val)
}
