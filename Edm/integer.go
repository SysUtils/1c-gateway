package Edm

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Integer struct {
	int
}

func NewInteger(int int) *Integer {
	return &Integer{int: int}
}

func (i Integer) MarshalJSON() ([]byte, error) {
	val := strconv.Itoa(i.int)

	return json.Marshal(val)
}

func (i *Integer) UnmarshalJSON(buf []byte) error {
	val, err := strconv.Atoi(strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	i.int = val
	return nil
}

func (i *Integer) Value() int {
	return i.int
}
