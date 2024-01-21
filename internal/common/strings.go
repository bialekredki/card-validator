package common

import (
	"errors"
	"strconv"
)

func getNthDigit(input string, n int) (int, error) {
	if len(input)-1 < n {
		return 0, errors.New("out of bounds")
	}

	return strconv.Atoi(string(input[n]))
}

func LastDigit(input string) (int, error) {
	return getNthDigit(input, len(input)-1)
}

func FirstDigit(input string) (int, error) {
	return getNthDigit(input, 0)
}
