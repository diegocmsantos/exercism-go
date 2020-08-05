package romannumerals

import (
	"errors"
	"strings"
)

var romanSymbols = map[int]string{
	1: "I",
	5: "V",
	10: "X",
	50: "L",
	100: "C",
	500: "D",
	1000: "M",
}

// ToRomanNumeral receives an arabic number and translate it to a roman number representation.
func ToRomanNumeral(arabic int) (string, error) {

	if arabic > 3000 {
		return "", errors.New("arabic number greater than 3000")
	}

	if arabic < 0 {
		return "", errors.New("cannot convert negative numbers")
	}

	var romanNumber string
	var quotient int
	var rest = arabic
	if rest >= 1000 {
		quotient += rest / 1000
		rest = rest % 1000
		romanNumber += strings.Repeat(romanSymbols[1000], quotient)
	}
	if rest >= 500 {
		quotient += rest / 500
		rest = rest % 500
		romanNumber += strings.Repeat(romanSymbols[500], quotient)
	}
	if rest >= 100 {
		quotient += rest / 100
		rest = rest % 100
		romanNumber += strings.Repeat(romanSymbols[100], quotient)
	}
	if rest >= 10 {
		quotient += rest / 10
		rest = rest % 10
		romanNumber += strings.Repeat(romanSymbols[10], quotient)
	}
	if rest >= 4 && rest < 9 {
		quotient = rest / 5
		if quotient == 0 {
			romanNumber += romanSymbols[1] + romanSymbols[5]
		} else {
			rest = rest % 5
			if rest == 0 {
				romanNumber += romanSymbols[5]
			} else {
				romanNumber += strings.Repeat(romanSymbols[5], rest)
			}
		}
	}
	if rest < 4 {
		romanNumber += strings.Repeat(romanSymbols[1], rest)
	}

	return romanNumber, nil
}
