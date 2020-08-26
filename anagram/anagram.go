package anagram

import (
  "strings"
  "sync"
)

func Detect(subject string, candidates []string) []string {
  var found []string
  output := make(chan string, len(candidates))

  var wg sync.WaitGroup
  wg.Add(len(candidates))
  for _, c := range candidates {
    go isAnagram(subject, c, output, &wg)
  }

  wg.Wait()
  close(output)
  for s := range output {
    if s != "" {
      found = append(found, s)
    }
  }

  return found
}

func isAnagram(subject, candidate string, output chan string, wg *sync.WaitGroup) {
  defer wg.Done()
  subject = strings.ToLower(subject)
  newCandidate := strings.ToLower(candidate)

  if len(subject) != len(newCandidate) || subject == newCandidate {
    output <- ""
    return
  }

  runeArray := []rune(subject)
  for _, l := range runeArray {
    candidateLenBefore := len(newCandidate)
    newCandidate = strings.Replace(newCandidate, string(l), "", 1)
    if candidateLenBefore == len(newCandidate) {
      output <- ""
      return
    }
  }
  output <- candidate
  return
}