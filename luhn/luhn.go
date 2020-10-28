package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if the given input is a valid number based on the Luhn formula
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	inputSize := len(input)
	if inputSize <= 1 {
		return false
	}
	var number, sum int
	double := false
	for i := inputSize - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(input[i])) {
			return false
		}
		number = int(input[i] - '0')
		if double {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		sum += number
		double = !double
	}
	return sum%10 == 0
}
