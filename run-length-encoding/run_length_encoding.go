package encode

import (
  "regexp"
  "strconv"
  "strings"
)

// RunLengthEncode encodes a given string
func RunLengthEncode(input string) string {
  if len(input) == 0 {
    return ""
  }

  var result string
  tempLetter := input[0]
  counter := 1
  for i := 1; i < len(input); i++ {
    if input[i] == tempLetter {
      counter++
    } else {
      result = buildResult(counter, result, tempLetter)
      tempLetter = input[i]
      counter = 1
    }
  }
  return buildResult(counter, result, tempLetter)
}

func buildResult(counter int, result string, tempLetter uint8) string {
  if counter == 1 {
    result += string(tempLetter)
  } else {
    result += strconv.Itoa(counter) + string(tempLetter)
  }
  return result
}

const numberPlusLetterRegex = `\d*(\w|\s)`
const numberRegex = `\d+`

// RunLengthDecode decodes a given string
func RunLengthDecode(input string) string {
  var result string
  re := regexp.MustCompile(numberPlusLetterRegex)
  reOnlyNumber := regexp.MustCompile(numberRegex)

  for _, c := range re.FindAllString(input, -1) {
    unit := reOnlyNumber.FindAllString(c, -1)
    if len(unit) > 0 {
      number, _ := strconv.Atoi(unit[0])
      letter := strings.ReplaceAll(c, unit[0], "")
      result += explodeString(number, letter)
    } else {
      result += c
    }
  }

  if result == "" {
    return input
  }
  return result
}

func explodeString(number int, letter string) string {
  var result string
  for i := 0; i < number; i++ {
    result += letter
  }
  return result
}