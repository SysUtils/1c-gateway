package schema

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
)

var subscriptionValidPrefixes = []string{
	"AccumulationRegister_",
	"InformationRegister_",
}
var subscriptionValidSuffixes = []string{
	"_RecordType",
	"Turnover",
}

func (g *Generator) genSubscriptions(source []shared.OneCType) string {
	queries := "type Subscription {\n"
	for _, entity := range source {
		validPrefix := false
		validSuffix := false
		for _, s := range subscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}

		for _, s := range subscriptionValidSuffixes {
			if strings.HasSuffix(entity.Name, s) {
				validSuffix = true
				break
			}
		}
		if validPrefix && validSuffix {
			queries += g.genSubscription(entity)
			queries += "\n"
		}
	}
	queries += fmt.Sprintf(`}`)
	return queries
}

func (g *Generator) genSubscription(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`	OnCreate%s: %s!`, t, t)
}
