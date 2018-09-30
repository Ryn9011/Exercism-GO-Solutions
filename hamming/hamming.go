package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {

	hammingDistance := 0

	err := errors.New("Sequences must be of the same length.")

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hammingDistance++
			}
		}
		return hammingDistance, nil
	}
	return 0, err
}
