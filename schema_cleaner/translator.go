package schema_cleaner

import (
	"strings"
)

func TranslateType(src string) string {
	if strings.HasPrefix(src, "Edm.") {
		src = src[4:]
	}

	if strings.HasPrefix(src, "StandardODATA.") {
		src = src[14:]
	}
	if strings.HasPrefix(src, "Collection(") && strings.HasSuffix(src, ")") {
		src = TranslateType(src[11 : len(src)-1])
	}
	return src
}
