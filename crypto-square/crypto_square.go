package cryptosquare

import (
	"math"
	"strings"
)

type Block struct {
	block [100]string
}

func Encode(pt string) string {
	normalizedInput := Normalize(pt)
	col, row := CalculateRectangleDimensions(len(normalizedInput))
	c := int(col)
	r := int(row)
	rectangleDimensions := int(c * r)
	rectangle := make([]Block, rectangleDimensions)
	normInputSize := len(normalizedInput)

	// Create blocks and add them to rectangle
	var block Block
	count := 0
	for i := 0; i < normInputSize; i++ {
		for j := 0; j < c && (count < normInputSize+((c*int(r))-normInputSize)); j++ {
			if count < len(normalizedInput) {
				block.block[j] = string(normalizedInput[count])
				count++
			} else {
				block.block[j] = " "
				count++
			}
		}
		rectangle[i] = block
	}

	// Create coded string from columns of rectangle
	codedString := ""
	blockIndex := 0
	for i := 0; i < c; i++ {
		for j := 0; j < int(r); j++ {
			codedString += rectangle[j].block[blockIndex]
		}
		blockIndex++
		if i < c-1 {
			codedString += " "
		}
	}
	return codedString
}

func CalculateRectangleDimensions(n int) (c float64, r float64) {
	floatOfNumber := float64(n)
	var squareOfFloat = math.Sqrt(floatOfNumber)
	squareOfFloatRounded := math.Round(squareOfFloat)
	temp := math.Round(squareOfFloat)

	if squareOfFloatRounded*squareOfFloatRounded == floatOfNumber {
		c = squareOfFloat
		r = squareOfFloat
		return
	}

	if temp > squareOfFloat {
		r = temp
		if (r * r) >= floatOfNumber {
			c = r
		} else {
			c = r + 1
		}
		return
	} else {
		r = temp
		c = r + 1
		return
	}

}

func Normalize(s string) string {
	s = strings.ToLower(s)
	norm := ""

	for i := 0; i < len(s); i++ {
		l := s[i : i+1]
		if (l >= "a" && l <= "z") || (l >= "0" && l <= "9") {
			norm += l
		}
	}
	return norm
}
