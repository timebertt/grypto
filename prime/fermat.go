package prime

import (
  "math/rand"

  "github.com/timebertt/grypto/modular"
)

// IsPrimeFermat32 tests p for primality using Fermat's primality test.
// It repeatedly chooses a random integer between 2 and p-2 and tests if a^(p-1) ≡ 1 mod p (rounds times).
// If p is prime, the condition must be true (according to Fermat's little theorem).
// That means, that if a^(p-1) is not congruent to 1 mod p, p is definitely not prime and the test stops.
// If the test doesn't find any integer contradicting Fermat's little theorem during the given number of rounds,
// it returns true, meaning that p is pseudo prime, but we can't actually be sure if p is really prime.
// There is an infinite number of integers (called "Carmichael numbers") for which Fermat's little theorem will be
// true for all integers (all integers will be "Fermat liars"). That's a serious flaw in Fermat's primality test.
// That's why we should rather use other probabilistic primality tests (e.g. Miller-Rabin primality test), for
// for which we can at least tell the probability of error.
// See: https://en.wikipedia.org/wiki/Fermat_primality_test
func IsPrimeFermat32(p int32, rounds int) (isPrime bool) {
  if p <= 1 {
    // 1 is neither prime nor composite
    return false
  }

  // handle simple cases
  if p <= 3 {
    return true
  }
  if p%2 == 0 {
    // if p is even we can directly say, that p is not prime
    return false
  }

  for r := 0; r < rounds; r++ {
    // choose random integer in range [2,n-2]
    a := rand.Int31n(p-3) + 2

    // a^(p-1) ≡ 1 mod p must be true if p is prime
    if modular.Pow32(a, p-1, p) != 1 {
      // p is definitely not prime
      return false
    }
  }

  // p is pseudo prime, we can't be sure if p is really prime
  return true
}
