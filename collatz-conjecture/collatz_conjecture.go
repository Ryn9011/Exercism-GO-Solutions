package collatzconjecture

import "errors"

func CollatzConjecture(num int) (i int, err error) {

	error := errors.New("Integers must be greater than zero")

	if num <= 0 {
		return 0, error
	}
	for i = 0; num != 1; i++ {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
	}
	return
}
