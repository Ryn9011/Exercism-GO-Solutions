package etl

import (
	"strings"
)

func Transform(legacy map[int][]string) map[string]int {
	modern := map[string]int{}

	for key, letters := range legacy {
		for _, letter := range letters {
			modern[strings.ToLower(letter)] = key
		}
	}
	return modern
}
