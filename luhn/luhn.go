package luhn

import (
	"strconv"
	"strings"
)

func Valid(numbers string) bool {
	numbers = strings.Replace(numbers, " ", "", -1)
	var sum int

	if _, err := strconv.Atoi(numbers); err != nil {
		return false
	}

	for i := 0; i < len(numbers); i++ {

		digit, _ := strconv.Atoi(numbers[i : i+1])

		if len(numbers) == 1 && digit == 0 {
			return false
		}

		if (len(numbers)-i)%2 == 0 {
			double := digit * 2
			if double > 9 {
				double -= 9
			}
			sum += double
		} else {
			sum += digit
		}

	}
	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}
