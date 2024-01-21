package algorithm

import (
	"bialekredki/card-validator/internal/common"
	"errors"
	"strconv"
)

const invalidCardNumberErrorMessage = "invalid card numer"

var invalidCardNumberError = errors.New(invalidCardNumberErrorMessage)

func reverseCardNumber(input string) []rune {
	inputRunes := []rune(input)
	var result = make([]rune, 0)
	for i := len(inputRunes) - 1; i >= 0; i-- {
		result = append(result, inputRunes[i])
	}
	return result
}

func removeLastCharacter(input string) string {
	if len(input) == 0 {
		return input
	}
	return input[:len(input)-1]
}

func sumOfDigits(value int) int {
	sum := 0
	for value != 0 {
		sum += value % 10
		value = value / 10
	}
	return sum
}

func IsValid(cardNumber string) (bool, error) {
	reversedCardNumber := reverseCardNumber(removeLastCharacter(cardNumber))
	lastDigit, err := common.LastDigit(cardNumber)
	if err != nil {
		return false, invalidCardNumberError
	}
	sum := 0
	for i, r := range reversedCardNumber {
		value, err := strconv.Atoi(string(r))
		if err != nil {
			return false, invalidCardNumberError
		}
		if i%2 == 0 {
			value *= 2
			if value > 9 {
				value = sumOfDigits(value)
			}
		}
		sum += value
	}

	return 10-(sum%10) == lastDigit, nil
}
