package prime_test

import (
  "math"
  "testing"

  . "github.com/onsi/ginkgo"

  "github.com/timebertt/grypto/prime"
)

var _ = Describe("IsPrimeMillerRabin32", func() {
  It("should correctly detect primes", func() {
    testPrimes1000(func(i int32) bool {
      return prime.IsPrimeMillerRabin32(i, 12)
    })
  })
})

func BenchmarkIsPrimeMillerRabin32(b *testing.B) {
  for i := 0; i < b.N; i++ {
    prime.IsPrimeMillerRabin32(math.MaxInt32, 12)
  }
}
