package lsproduct

import (
  "errors"
  "strconv"
  "sync"
)

// LargestSeriesProduct calculates the product of all digits within a span and returns the greatest one.
func LargestSeriesProduct(digits string, span int) (int64, error) {

  if span < 0 {
    return -1, errors.New("span must be greater than zero")
  }

  if span > len(digits) {
    return -1, errors.New("span must be smaller than string length")
  }

  if digits == "" && span == 0 {
    return 1, nil
  }

  var result int64
  output := make(chan int64, len(digits))
  var wg sync.WaitGroup
  for i:= 0; i < len(digits); i++ {
    if i+span-1 < len(digits) {
      wg.Add(1)
      go product(digits[i:i+span], output, &wg)
    }
  }
  wg.Wait()
  close(output)
  for d := range output {
    if d == -1 {
      return d, errors.New("digits input must only contain digits")
    }
    if result < d {
      result = d
    }
  }

  return result, nil
}

func product(digits string, ch chan int64, wg *sync.WaitGroup) {
  defer wg.Done()
  var result int64 = 1
  for _, l := range digits {
    n, err := strconv.ParseInt(string(l), 10, 64)
    if err != nil {
      ch <- -1
      return
    }
    result = result * n
  }
  ch <- result
}