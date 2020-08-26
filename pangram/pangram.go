package pangram

import "unicode"

const alphabetSize = 25

func IsPangram(input string) bool {
  if len(input) == 0 {
    return false
  }
  lettersMap := make(map[rune]bool)
  for _, l := range input {
    lowerCaseRune := unicode.ToLower(l)
    if ok := lettersMap[lowerCaseRune]; ok || !unicode.IsLetter(lowerCaseRune) {
      continue
    }
    lettersMap[lowerCaseRune] = true
  }
  return len(lettersMap) > alphabetSize
}