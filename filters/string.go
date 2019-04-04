package filters

import "fmt"

type StringOp int

const (
	StringOpEq StringOp = iota
	StringOpNe
	StringOpSubstringof
	StringOpEndswith
	StringOpStartswith
)

func String(field string, op StringOp, value string) *Filter {
	switch op {
	case StringOpEq:
		return &Filter{fmt.Sprintf("%s eq '%s'", field, EscapeString(value))}
	case StringOpNe:
		return &Filter{fmt.Sprintf("%s ne '%s'", field, EscapeString(value))}
	case StringOpSubstringof:
		return &Filter{fmt.Sprintf("substringof(%s, '%s')", field, EscapeString(value))}
	case StringOpEndswith:
		return &Filter{fmt.Sprintf("endswith(%s, '%s')", field, EscapeString(value))}
	case StringOpStartswith:
		return &Filter{fmt.Sprintf("startswith(%s, '%s')", field, EscapeString(value))}
	}
	return nil
}
