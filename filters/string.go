package filters

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/static"
)

type StringOp int

const (
	StringOpEq StringOp = iota
	StringOpNe
	StringOpSubstringof
	StringOpEndswith
	StringOpStartswith
)

func String(field string, op StringOp, value string) *odata.Filter {
	switch op {
	case StringOpEq:
		return &odata.Filter{fmt.Sprintf("%s eq '%s'", field, odata.EscapeString(value))}
	case StringOpNe:
		return &odata.Filter{fmt.Sprintf("%s ne '%s'", field, odata.EscapeString(value))}
	case StringOpSubstringof:
		return &odata.Filter{fmt.Sprintf("substringof(%s, '%s')", field, odata.EscapeString(value))}
	case StringOpEndswith:
		return &odata.Filter{fmt.Sprintf("endswith(%s, '%s')", field, odata.EscapeString(value))}
	case StringOpStartswith:
		return &odata.Filter{fmt.Sprintf("startswith(%s, '%s')", field, odata.EscapeString(value))}
	}
	return nil
}
