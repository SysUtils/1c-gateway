package odata

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Guid string

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Guid) ImplementsGraphQLType(name string) bool {
	return name == "Guid"
}

// UnmarshalGraphQL is a custom unmarshaler for Time
//
// This function will be called whenever you use the
// time scalar as an input

func (t *Guid) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = Guid(strings.Trim(input, `"`))
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t Guid) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *Guid) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if strings.HasPrefix(s, "guid'") {
		s = s[5 : len(s)-1]
	}
	*t = Guid(s)
	return nil
}

func (t Guid) AsParameter() string {
	return `guid'` + string(t) + `'`
}
