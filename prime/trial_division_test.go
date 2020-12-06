package prime_test

import (
  "math"
  "testing"

  . "github.com/onsi/ginkgo"

  "github.com/timebertt/grypto/prime"
)

var _ = Describe("IsPrimeTrialDivision32", func() {
  It("should correctly detect primes", func() {
    testPrimes1000(prime.IsPrimeTrialDivision32)
  })
})

func BenchmarkIsPrimeTrialDivision32(b *testing.B) {
  for i := 0; i < b.N; i++ {
    prime.IsPrimeTrialDivision32(math.MaxInt32)
  }
}
