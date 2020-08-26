package wordcount

import (
  "regexp"
  "strings"
)

type Frequency map[string]int

const onlyWordsAndNumbers = `(\w)+'?(\w)+|\w`
const specialCharacters = `[\t\n]`

func WordCount(phrase string) Frequency {
  var freqMap = make(Frequency)
  var re = regexp.MustCompile(specialCharacters)
  phraseWithNoSpecialCharacters := re.ReplaceAllString(phrase, "")
  r := regexp.MustCompile(onlyWordsAndNumbers)
  wordsArray := r.FindAllString(phraseWithNoSpecialCharacters, -1)
  for _, word := range wordsArray {
    freqMap[strings.ToLower(word)]++
  }
  return freqMap
}
