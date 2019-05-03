package odata

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Binary []byte

func (Binary) ImplementsGraphQLType(name string) bool {
	return name == "Binary"
}

func (t *Binary) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		res, err := b64.StdEncoding.DecodeString(input)
		*t = Binary(res)
		return err
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(b64.StdEncoding.EncodeToString(t))
}

func (t *Binary) UnmarshalJSON(b []byte) error {
	val, err := b64.StdEncoding.DecodeString(strings.Trim(string(b), `"`))
	*t = Binary(val)
	return err
}

func (t Binary) AsParameter() string {
	return `'` + b64.StdEncoding.EncodeToString(t) + `'`
}
