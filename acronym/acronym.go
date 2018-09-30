package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var b strings.Builder

	fields := strings.FieldsFunc(s, Delimeter)
	for _, field := range fields {
		b.WriteString(string([]rune(field)[0]))
	}

	return strings.ToUpper(b.String())
}

func Delimeter(c rune) bool {
	return unicode.IsSpace(c) || c == '-'
}
