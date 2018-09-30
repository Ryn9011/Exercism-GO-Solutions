package twofer

import (
	"fmt"
)

func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprint("One for ", name, ", one for me.")
}
