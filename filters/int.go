package filters

import "fmt"

type IntOp int

const (
	IntOpEq IntOp = iota
	IntOpNe
	IntOpGt
	IntOpGe
	IntOpLt
	IntOpLe
)

func Int(field string, op IntOp, value int) *Filter {
	switch op {
	case IntOpEq:
		return &Filter{fmt.Sprintf("%s eq %d", field, value)}
	case IntOpNe:
		return &Filter{fmt.Sprintf("%s ne %d", field, value)}
	case IntOpGt:
		return &Filter{fmt.Sprintf("%s gt %d", field, value)}
	case IntOpGe:
		return &Filter{fmt.Sprintf("%s ge %d", field, value)}
	case IntOpLt:
		return &Filter{fmt.Sprintf("%s lt %d", field, value)}
	case IntOpLe:
		return &Filter{fmt.Sprintf("%s le %d", field, value)}
	}
	return nil
}
