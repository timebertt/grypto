package modular_test

import (
  "math"
  "testing"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/modular"
)

var _ = Describe("Pow32", func() {
  It("should panic on invalid inputs", func() {
    test := func(b, e, m int32) {
      ExpectWithOffset(1, func() {
        modular.Pow32(b, e, m)
      }).To(Panic())
    }

    test(1, 2, -1)
    test(1, 2, 0)
    test(1, -1, 1)
    test(0, 0, 1)
  })

  It("should correctly calculate modular exponentiation", func() {
    test := func(b, e, m, expected int32) {
      ExpectWithOffset(1, modular.Pow32(b, e, m)).To(Equal(expected))
    }

    // exp = 0
    test(5, 0, 7, 1)
    test(-5, 0, 7, 1)

    // exp > 0
    test(2, 35, 561, 263)
    test(34, 560, 561, 34)
    test(180, 15, 23, 20)
    test(180, 11, 17, 3)

    // base < 0
    test(-2, 35, 561, 298)

    // base = 0
    test(0, 35, 561, 0)

    // int32 overflows
    test(math.MaxInt32-3, math.MaxInt32, math.MaxInt32-1, 2147483518)
  })
})

var result int32

func BenchmarkPow32MaxInt32(b *testing.B) {
  var r int32
  for n := 0; n < b.N; n++ {
    r = modular.Pow32(math.MaxInt32-3, math.MaxInt32, math.MaxInt32-1)
  }
  result = r
}
