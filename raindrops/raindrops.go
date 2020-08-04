package raindrops

import (
	"strconv"
)

// Convert takes a number and:
// if has 3 as factor adds "Pling" to the result
// if has 5 as factor adds "Plang" to the result
// if has 7 as factor adds "Plong" to the result
// if it doesn't have 3, 5 or 7 as factors adds the number itself to the result
func Convert(number int) string {
	result := ""
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		result += strconv.Itoa(number)
	}

	return result
}
