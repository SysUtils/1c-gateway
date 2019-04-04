package odata

import (
	"encoding/json"
	"fmt"
	"time"
)

type Time time.Time

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Time) ImplementsGraphQLType(name string) bool {
	return name == "Time"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *Time) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		tmp := Time(input)
		t = &tmp
		return nil
	case string:
		var err error
		time, err := time.Parse(time.RFC3339, input)
		tmp := Time(time)
		t = &tmp
		return err
	case int:
		time := time.Unix(int64(input), 0)
		tmp := Time(time)
		t = &tmp
		return nil
	case float64:
		time := time.Unix(int64(input), 0)
		tmp := Time(time)
		t = &tmp
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// MarshalJSON is a custom marshaler for Time
//
// This function will be called whenever you
// query for fields that use the Time type
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t))
}
