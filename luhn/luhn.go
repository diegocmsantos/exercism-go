package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if the given input is a valid number based on the Luhn formula
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	var number, sum int
	double := len(input)%2 == 0
	for _, r := range input {
		if !unicode.IsDigit(r) {
			return false
		}
		number = int(r - '0')
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
