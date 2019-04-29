package filters

import (
	"fmt"
	"strings"
)

type Filter struct {
	value string
}

func EscapeString(src string) string {
	return strings.Replace(src, "'", "''", -1)
}

func (f Filter) ToString() string {
	return fmt.Sprintf("(%s)", f.value)
}

func (f Filter) And(f1 *Filter) *Filter {
	return &Filter{f.ToString() + " and " + f1.ToString()}
}

func (f Filter) Or(f1 *Filter) *Filter {
	return &Filter{f.ToString() + " or " + f1.ToString()}
}

func (f Filter) AndNot(f1 *Filter) *Filter {
	return &Filter{f.ToString() + " and not " + f1.ToString()}
}
