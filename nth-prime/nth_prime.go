package prime

import (
  "math/big"
)

// Nth returns the nth prime number and true if found, 0 and false otherwise.
func Nth(n int) (int, bool) {
  if n == 0 {
    return 0, false
  }

  counter := 1
  for i := 2; true; i++ {
    if isPrime(i) {
      if counter == n {
        return i, true
      }
      counter++
    }
  }
  return 0, false
}

func isPrime(n int) bool {
  return big.NewInt(int64(n)).ProbablyPrime(0)
}