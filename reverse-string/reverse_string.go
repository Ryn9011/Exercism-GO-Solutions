package reverse

import (
	"strings"
)

func String(s string) string {
	var b strings.Builder
	runeSlice := []rune(s)

	for i := len(runeSlice) - 1; i >= 0; i-- {
		b.WriteString(string(runeSlice[i]))
	}
	return b.String()
}
