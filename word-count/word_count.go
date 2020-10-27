package wordcount

import (
  "regexp"
  "strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`[\t\n]`)
var r = regexp.MustCompile(`(\w)+'?(\w)+|\w`)

func WordCount(phrase string) Frequency {
  var freqMap = make(Frequency)
  phraseWithNoSpecialCharacters := re.ReplaceAllString(phrase, "")
  wordsArray := r.FindAllString(phraseWithNoSpecialCharacters, -1)
  for _, word := range wordsArray {
    freqMap[strings.ToLower(word)]++
  }
  return freqMap
}
