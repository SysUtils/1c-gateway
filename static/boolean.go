package odata

import (
	"encoding/json"
	"fmt"
)

type Boolean bool

func (Boolean) ImplementsGraphQLType(name string) bool {
	return name == "Boolean"
}

func (t *Boolean) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case bool:
		*t = Boolean(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(t))
}

func (t *Boolean) UnmarshalJSON(b []byte) error {
	*t = Boolean(string(b) == "true")
	return nil
}

func (t Boolean) AsParameter() string {
	if t {
		return "true"
	} else {
		return "false"
	}
}
