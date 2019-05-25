package filters

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/static"
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

func Date(field string, op DateOp, value time.Time) *odata.Filter {
	switch op {
	case DateOpEq:
		return &odata.Filter{fmt.Sprintf("%s eq datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpNe:
		return &odata.Filter{fmt.Sprintf("%s ne datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpGt:
		return &odata.Filter{fmt.Sprintf("%s gt datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpGe:
		return &odata.Filter{fmt.Sprintf("%s ge datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpLt:
		return &odata.Filter{fmt.Sprintf("%s lt datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	case DateOpLe:
		return &odata.Filter{fmt.Sprintf("%s le datetime'%s'", field, value.Format("2006-01-02T15:04:05"))}
	}
	return nil
}
