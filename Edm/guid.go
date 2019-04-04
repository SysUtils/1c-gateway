package Edm

import (
	"github.com/beevik/guid"
)

type Guid string

func NewGuid(data string) *Guid {
	res := Guid(data)
	return &res
}

func GenGuid() *Guid {
	res := Guid(guid.NewString())
	return &res
}

func (g Guid) IsNil() bool {
	return g == "00000000-0000-0000-0000-000000000000"
}
