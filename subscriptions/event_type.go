package subscriptions

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) genTypeConsts(createEntities, updateEntities []shared.OneCType) string {
	result := ""
	for _, val := range createEntities {
		result += fmt.Sprintf("\tCreate%s\n", g.translateType(val.Name))
	}
	for _, val := range updateEntities {
		result += fmt.Sprintf("\tUpdate%s\n", g.translateType(val.Name))
		result += fmt.Sprintf("\tDelete%s\n", g.translateType(val.Name))
	}
	return result
}
