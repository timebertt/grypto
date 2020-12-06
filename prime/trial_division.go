package prime

import (
  "math"
)

// IsPrimeTrialDivision32 tests p for primality using the trial division algorithm.
// It tests for every n = 2,3,... <= sqrt(p) if p can be divided by n. If it doesn't find any divisor of p,
// then p is prime. While this approach is very simple and easy to understand, it is very inefficient and impractical.
// In practice it is only used for testing small integers for primality (< 10^6).
// The same approach can also be used for factoring small integers.
func IsPrimeTrialDivision32(p int32) (isPrime bool) {
  if p <= 1 {
    // 1 is neither prime nor composite
    return false
  }

  for i := int32(2); i <= int32(math.Sqrt(float64(p))); i++ {
    if p%i == 0 {
      return false
    }
  }

  return true
}
