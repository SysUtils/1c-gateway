package odata

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

type Stream []byte

func (Stream) ImplementsGraphQLType(name string) bool {
	return name == "Stream"
}

func (t *Stream) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		res, _ := b64.StdEncoding.DecodeString(input)
		*t = Stream(res)
	default:
		return fmt.Errorf("wrong type")
	}
	return nil
}

func (t Stream) MarshalJSON() ([]byte, error) {
	return json.Marshal(b64.StdEncoding.EncodeToString(t))
}

func (t Stream) AsParameter() string {
	return `'` + b64.StdEncoding.EncodeToString(t) + `'`
}
