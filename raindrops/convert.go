package raindrops

import (
	"strconv"
	"strings"
)

func Convert(num int) string {
	b := strings.Builder{}

	rainSound := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		result := strconv.Itoa(num)
		return result
	}

	if num%3 == 0 {
		b.WriteString(rainSound[3])
	}
	if num%5 == 0 {
		b.WriteString(rainSound[5])
	}
	if num%7 == 0 {
		b.WriteString(rainSound[7])
	}
	return b.String()
}
