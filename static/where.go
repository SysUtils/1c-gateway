package odata

import (
	"fmt"
	"net/url"
	"strings"
)

// Interface for odata Filter type
type IFilter interface {
	ToString() string
}

type Where struct {
	Top     int32
	Skip    int32
	Filter  IFilter
	Orderby string
	Fields  []string
}

// Returns string representation of Where object
func (w *Where) Serialize() string {
	params := ""

	if w.Filter != nil {
		params += fmt.Sprintf("$filter=%s&", url.PathEscape((w.Filter).ToString()))
	}

	if w.Orderby != "" {
		params += fmt.Sprintf("$orderby=%s&", url.PathEscape(w.Orderby))
	}

	if w.Top != 0 {
		params += fmt.Sprintf("$top=%d&", w.Top)
	}

	if w.Skip != 0 {
		params += fmt.Sprintf("$skip=%d&", w.Skip)
	}

	if w.Fields != nil {
		params += fmt.Sprintf("$select=%s&", url.PathEscape(strings.Join(w.Fields, ", ")))
	}
	return params
}
