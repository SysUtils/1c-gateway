package odata

import (
	"encoding/json"
	"fmt"
	"strings"
)

type String string

func (String) ImplementsGraphQLType(name string) bool {
	return name == "String"
}

func (t *String) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*t = String(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (t String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *String) UnmarshalJSON(b []byte) error {
	s := ""
	err := json.Unmarshal(b, &s)
	*t = String(s)
	return err
}

func (t String) Escape() string {
	return strings.Replace(string(t), "'", "''", -1)
}

func (t String) AsParameter() string {
	return fmt.Sprintf("'%s'", t.Escape())
}
