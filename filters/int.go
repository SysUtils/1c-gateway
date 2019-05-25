package filters

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/static"
)

type IntOp int

const (
	IntOpEq IntOp = iota
	IntOpNe
	IntOpGt
	IntOpGe
	IntOpLt
	IntOpLe
)

func Int(field string, op IntOp, value int) *odata.Filter {
	switch op {
	case IntOpEq:
		return &odata.Filter{fmt.Sprintf("%s eq %d", field, value)}
	case IntOpNe:
		return &odata.Filter{fmt.Sprintf("%s ne %d", field, value)}
	case IntOpGt:
		return &odata.Filter{fmt.Sprintf("%s gt %d", field, value)}
	case IntOpGe:
		return &odata.Filter{fmt.Sprintf("%s ge %d", field, value)}
	case IntOpLt:
		return &odata.Filter{fmt.Sprintf("%s lt %d", field, value)}
	case IntOpLe:
		return &odata.Filter{fmt.Sprintf("%s le %d", field, value)}
	}
	return nil
}
