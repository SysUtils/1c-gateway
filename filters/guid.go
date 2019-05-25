package filters

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/static"
)

type GuidOp int

const (
	GuidOpEq GuidOp = iota
	GuidOpNe
)

func Guid(field string, op GuidOp, value odata.Guid) *odata.Filter {
	switch op {
	case GuidOpEq:
		return &odata.Filter{fmt.Sprintf("%s eq guid'%s'", field, value)}
	case GuidOpNe:
		return &odata.Filter{fmt.Sprintf("%s ne guid'%s'", field, value)}
	}
	return nil
}
