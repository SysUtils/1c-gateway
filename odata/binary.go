package odata

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

type Binary []byte

func (Binary) ImplementsGraphQLType(name string) bool {
	return name == "Binary"
}

func (t *Binary) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		res, _ := b64.StdEncoding.DecodeString(input)
		*t = Binary(res)
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(b64.StdEncoding.EncodeToString(t))
}

func (t Binary) AsParameter() string {
	return `'` + b64.StdEncoding.EncodeToString(t) + `'`
}
