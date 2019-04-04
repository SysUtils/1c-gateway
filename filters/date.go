package filters

import (
	"fmt"
	"time"
)

type DateOp int

const (
	DateOpEq DateOp = iota
	DateOpNe
	DateOpGt
	DateOpGe
	DateOpLt
	DateOpLe
)

func Date(field string, op DateOp, value time.Time) *Filter {
	switch op {
	case DateOpEq:
		return &Filter{fmt.Sprintf("%s eq datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpNe:
		return &Filter{fmt.Sprintf("%s ne datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpGt:
		return &Filter{fmt.Sprintf("%s gt datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpGe:
		return &Filter{fmt.Sprintf("%s ge datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpLt:
		return &Filter{fmt.Sprintf("%s lt datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpLe:
		return &Filter{fmt.Sprintf("%s le datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	}
	return nil
}
