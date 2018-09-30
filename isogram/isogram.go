package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {

	word := strings.ToLower(input)

	runes := map[rune]bool{}

	for _, r := range word {
		if runes[r] && unicode.IsLetter(r) {
			return false
		}
		runes[r] = true
	}
	return true
}
