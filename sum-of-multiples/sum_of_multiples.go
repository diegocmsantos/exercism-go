package summultiples

import (
  "sync"
)

func SumMultiples(limit int, divisors... int) int {
  set := make(map[int]bool)
  intChan := make(chan []int, len(divisors))
  var wg sync.WaitGroup
  wg.Add(len(divisors))
  for _, d := range divisors {
    go findMultiples(limit, d, intChan, &wg)
  }
  wg.Wait()
  close(intChan)
  for ii := range intChan {
    for _, i := range ii {
      set[i] = true
    }
  }

  var result int
  for k, _ := range set {
    result += k
  }

  return result
}

func findMultiples(limit, divisor int, ch chan []int, wg *sync.WaitGroup) {
  defer wg.Done()
  var number int
  var result []int

  if divisor == 0 {
    ch <- result
    return
  }

  for ok := true; ok; ok = number < limit {
    if number % divisor == 0 {
      result = append(result, number)
    }
    number++
  }
  ch <- result
}