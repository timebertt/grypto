package prime

import (
  "math/rand"

  "github.com/timebertt/grypto/modular"
)

// IsPrime32 tests p for primality using the default test (Miller-Rabin with 12 rounds).
func IsPrime32(p int32) bool {
  return IsPrimeMillerRabin32(p, 12)
}

// IsPrimeMillerRabin32 tests p for primality using the Miller-Rabin primality test.
// It repeatedly chooses a random integer between 2 and p-2 and tests if a set of equalities hold true for p with
// a given a.
// If neither of the equalities hold true, then p is composite and the test will return false.
// If the test doesn't find any integer for which neither of the equalities are met during the given number of rounds,
// it returns true, meaning that p is probably prime, but we can't actually be sure if p is really prime.
// The advantage of this test is, that we can at least tell the probability of a falsy reported prime and we can
// decrease this probability by increasing the number of rounds.
// If p is composite, then doing m rounds will report p as "probably prime" with a probability of 4^-m.
// See: https://en.wikipedia.org/wiki/Miller-Rabin_primality_test
func IsPrimeMillerRabin32(p int32, rounds int) (isPrime bool) {
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

  // write p-1 as 2^s * d with d odd, by factoring out powers of 2
  var (
    s = int32(0)
    d = p - 1
  )
  for d%2 == 0 {
    s++
    d /= 2
  }

outer:
  for r := 0; r < rounds; r++ {
    // choose random integer in range [2,n-2]
    a := rand.Int31n(p-3) + 2

    ad := modular.Pow32(a, d, p)
    if ad == 1 || ad == p-1 {
      // a is not a witness against p's primality, choose next a
      continue
    }
    for r := int32(0); r < s; r++ {
      ad = modular.Pow32(ad, 2, p)
      if ad == p-1 {
        // a is not a witness against p's primality, choose next a
        continue outer
      }
    }

    // a is witness against p's primality, p is definitely composite
    return false
  }

  // p is probably prime, for all chosen integers
  return true
}
