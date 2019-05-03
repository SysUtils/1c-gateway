package filters

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/static"
)

type BoolOp int

const (
	BoolOpEq BoolOp = iota
	BoolOpNe
)

func Bool(field string, op BoolOp, value bool) *odata.Filter {
	switch op {
	case BoolOpEq:
		if value {
			return &odata.Filter{fmt.Sprintf("%s eq true", field)}
		} else {
			return &odata.Filter{fmt.Sprintf("%s eq false", field)}
		}
	case BoolOpNe:
		if value {
			return &odata.Filter{fmt.Sprintf("%s ne true", field)}
		} else {
			return &odata.Filter{fmt.Sprintf("%s ne false", field)}
		}
	}
	return nil
}
