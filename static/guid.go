package odata

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Type for Edm.Guid
type Guid string

// Maps Guid to the graphql scalar type in the schema.
func (Guid) ImplementsGraphQLType(name string) bool {
	return name == "Guid"
}

// A custom unmarshaler for Guid type
func (t *Guid) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = Guid(strings.Trim(input, `"`))
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

// A custom json/graphql marshaller for Guid type
func (t Guid) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

// A custom json unmarshaller for Guid type
func (t *Guid) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if strings.HasPrefix(s, "guid'") {
		s = s[5 : len(s)-1]
	}
	*t = Guid(s)
	return nil
}

// A custom marshaller to uri query format for Guid type
func (t Guid) AsParameter() string {
	return `guid'` + string(t) + `'`
}
