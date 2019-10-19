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
var updateSubscriptionValidPrefixes = []string{
	"Catalog_",
	"Document_",
}

func (g *Generator) genSubscriptions(source []shared.OneCType) string {
	queries := "type Subscription {\n"
	for _, entity := range source {
		validPrefix := false
		validSuffix := false
		lite := false
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
			queries += g.genCreateSubscription(entity)
			queries += "\n"
			lite = true
		}

		validPrefix = false
		validSuffix = false
		for _, s := range updateSubscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}
		validSuffix = strings.Count(entity.Name, "_") == 1

		if validPrefix && validSuffix {
			queries += g.genUpdateSubscription(entity)
			queries += "\n"
			queries += g.genDeleteSubscription(entity)
			queries += "\n"
			if !lite {
				queries += g.genCreateSubscription(entity)
				queries += "\n"
			}
		}

	}
	queries += fmt.Sprintf(`}`)
	return queries
}

func (g *Generator) genCreateSubscription(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`	OnCreate%s: %s!`, t, t)
}

func (g *Generator) genUpdateSubscription(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`	OnUpdate%s: %s!`, t, t)
}

func (g *Generator) genDeleteSubscription(source shared.OneCType) string {
	t := g.translateType(source.Name)
	return fmt.Sprintf(
		`	OnDelete%s: %s!`, t, t)
}
