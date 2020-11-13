package romannumerals

import (
	"errors"
	"strings"
)

var romanSymbols = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

// ToRomanNumeral receives an arabic number and translate it to a roman number representation.
func ToRomanNumeral(arabic int) (string, error) {

	if arabic > 3000 {
		return "", errors.New("arabic number greater than 3000")
	}

	if arabic <= 0 {
		return "", errors.New("cannot convert negative numbers")
	}

	var romanNumber string
	var quotient = arabic
	var rest int
	for i := 1; quotient > 0; i = i * 10 {
		quotient = arabic / (i * 10)
		rest = arabic % (i * 10)
		romanNumber = translate(rest, i) + romanNumber
	}

	return romanNumber, nil
}

func translate(rest, decimal int) string {
	rest = rest / decimal
	romanNumber := ""
	if rest > 0 && rest <= 3 {
		romanNumber = strings.Repeat(romanSymbols[decimal], rest)
	} else if rest == 4 {
		romanNumber = romanSymbols[4*decimal]
	} else if rest > 4 && rest < 9 {
		romanNumber = romanSymbols[5*decimal] + strings.Repeat(romanSymbols[decimal], rest%5)
	} else if rest == 9 {
		romanNumber = romanSymbols[9*decimal]
	}
	return romanNumber
}
