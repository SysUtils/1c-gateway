package odata

import (
	"fmt"
	"net/url"
	"strings"
)

type IFilter interface {
	ToString() string
}

type Where struct {
	Top     int
	Skip    int
	Filter  IFilter
	Orderby string
	Fields  []string
}

func SerializeWhere(where Where) string {
	params := ""

	if where.Filter != nil {
		params += fmt.Sprintf("$filter=%s&", url.PathEscape((where.Filter).ToString()))
	}

	if where.Orderby != "" {
		params += fmt.Sprintf("$orderby=%s&", url.PathEscape(where.Orderby))
	}

	if where.Top != 0 {
		params += fmt.Sprintf("$top=%d&", where.Top)
	}

	if where.Skip != 0 {
		params += fmt.Sprintf("$skip=%d&", where.Skip)
	}

	if where.Fields != nil {
		params += fmt.Sprintf("$select=%s&", url.PathEscape(strings.Join(where.Fields, ", ")))
	}
	return params
}
