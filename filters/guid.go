package filters

import (
	"fmt"
	"gitlab.com/zullpro/core/1cclientgenerator.git/odata"
)

type GuidOp int

const (
	GuidOpEq GuidOp = iota
	GuidOpNe
)

func Guid(field string, op GuidOp, value odata.Guid) *Filter {
	switch op {
	case GuidOpEq:
		return &Filter{fmt.Sprintf("%s eq guid'%s'", field, value)}
	case GuidOpNe:
		return &Filter{fmt.Sprintf("%s ne guid'%s'", field, value)}
	}
	return nil
}
