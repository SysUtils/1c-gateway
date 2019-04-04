package filters

import (
	"fmt"
)

type BoolOp int

const (
	BoolOpEq BoolOp = iota
	BoolOpNe
)

func Bool(field string, op BoolOp, value bool) *Filter {
	switch op {
	case BoolOpEq:
		if value {
			return &Filter{fmt.Sprintf("%s eq true", field)}
		} else {
			return &Filter{fmt.Sprintf("%s eq false", field)}
		}
	case BoolOpNe:
		if value {
			return &Filter{fmt.Sprintf("%s ne true", field)}
		} else {
			return &Filter{fmt.Sprintf("%s ne false", field)}
		}
	}
	return nil
}
