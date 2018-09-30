package grains

import (
	"errors"
)

func Square(num int) (uint64, error) {
	var board [64]int
	board[0] = 1
	var returnVal uint64
	err := errors.New("Input not valid")

	if num <= 0 || num > 64 {
		return 0, err
	}
	if num == 1 {
		return 1, nil
	}

	for i := 1; i < num; i++ {
		board[i] = board[i-1] * 2
		if uint64(board[i]) > returnVal {
			returnVal = uint64(board[i])
		}
	}
	return returnVal, nil

}

func Total() uint64 {
	var board [64]int
	board[0] = 1
	total := 0

	for i := 1; i < 64; i++ {
		board[i] = board[i-1] * 2
		total += board[i]
	}
	return uint64(total) + 1
}
