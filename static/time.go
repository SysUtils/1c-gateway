package odata

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Type for Edm.Time
type Time time.Time

// Maps Time to the graphql scalar type in the schema.
func (Time) ImplementsGraphQLType(name string) bool {
	return name == "Time"
}

// A custom unmarshaler for Time type
func (t *Time) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case time.Time:
		*t = Time(input)
		return nil
	case string:
		val, err := time.Parse(OneCTimeFormat, input)
		if err != nil {
			val, err = time.Parse(JSTimeFormat, input)
		}
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

// A custom json/graphql marshaller for Time type
func (t Time) MarshalJSON() ([]byte, error) {
	val := time.Time(t).Format(JSTimeFormat)
	return json.Marshal(val)
}

// A custom json unmarshaller for Time type
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	val, err := time.Parse(JSTimeFormat, s)
	*t = Time(val)
	return err
}

// A custom marshaller to uri query format for Time type
func (t Time) AsParameter() string {
	val := time.Time(t).Format(OneCTimeFormat)
	return fmt.Sprintf("datetime'%s'", val)
}
