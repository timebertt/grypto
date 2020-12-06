package prime_test

import (
  "math"
  "testing"

  . "github.com/onsi/ginkgo"

  "github.com/timebertt/grypto/prime"
)

var _ = Describe("IsPrimeFermat32", func() {
  It("should correctly detect primes", func() {
    testPrimes1000(func(i int32) bool {
      return prime.IsPrimeFermat32(i, 12)
    })
  })
})

func BenchmarkIsPrimeFermat32(b *testing.B) {
  for i := 0; i < b.N; i++ {
    prime.IsPrimeFermat32(math.MaxInt32, 12)
  }
}
