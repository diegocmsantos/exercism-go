package luhn

import (
  "errors"
  "strings"
  "unicode"
)
// Valid checks if the given input is a valid number based on the Luhn formula
func Valid(input string) bool {

  input = strings.ReplaceAll(input, " ", "")
  sum, err := doubleAndSum(input)
  if err != nil {
    return false
  }
  return sum%10 == 0
}

// doubleAndSum doubles every second number from right to left and returns the sum of all numbers.
func doubleAndSum(input string) (int, error) {

  if len(input) <= 1 {
    return 0, errors.New("input length is less or equal to one")
  }

  var sum, count int
  for i := len(input) - 1; i >= 0; i-- {
    if unicode.IsDigit(rune(input[i])) {
      number := checkCounterIsOddAndNumberIsGreaterThanNine(count, int(input[i]-'0'))
      sum += number
      count++
    } else {
      return 0, errors.New("invalid character")
    }
  }
  return sum, nil
}

func checkCounterIsOddAndNumberIsGreaterThanNine(count, number int) int {
  if count%2 != 0 {
    number = number * 2
    if number > 9 {
      number = number - 9
    }
  }
  return number
}
