package subscriptions

import (
	"github.com/SysUtils/1c-gateway/shared"
	"strings"
)

// returns createEntities, updateEntities
func filterEntities(source []shared.OneCType) ([]shared.OneCType, []shared.OneCType) {
	updateEntities := []shared.OneCType{}
	createEntities := []shared.OneCType{}
	for _, entity := range source {
		validPrefix := false
		validSuffix := false
		for _, s := range createSubscriptionValidPrefixes {
			if strings.HasPrefix(entity.Name, s) {
				validPrefix = true
				break
			}
		}

		for _, s := range updateSubscriptionValidSuffixes {
			if strings.HasSuffix(entity.Name, s) {
				validSuffix = true
				break
			}
		}
		if validPrefix && validSuffix {
			createEntities = append(createEntities, entity)
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
			updateEntities = append(updateEntities, entity)
		}
	}
	return createEntities, updateEntities
}
