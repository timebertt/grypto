package modular_test

import (
  "math"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/modular"
)

var _ = Describe("DLog32", func() {
  It("should panic on invalid inputs", func() {
    test := func(x, b, m int32) {
      ExpectWithOffset(1, func() {
        modular.DLog32(x, b, m)
      }).To(Panic())
    }

    test(0, 1, 2)
    test(-1, 1, 2)
    test(1, -1, 2)
    test(1, 1, -1)
    test(1, 0, 2)
    test(1, 1, 0)
    test(1, 2, math.MaxInt32)
  })

  It("should correctly calculate dlog", func() {
    test := func(x, b, m, expected int32, expectedExists bool) {
      s, ok := modular.DLog32(x, b, m)
      ExpectWithOffset(1, ok).To(Equal(expectedExists))
      if expectedExists {
        ExpectWithOffset(1, s).To(Equal(expected))
      }
    }

    test(1, 1, 13, 0, true)
    test(1, 2, 13, 0, true)
    test(3, 2, 13, 4, true)
    test(5, 2, 13, 9, true)
    test(9, 2, 13, 8, true)
    test(12, 2, 13, 6, true)
    test(11, 2, 13, 7, true)

    test(2, 1, 13, 0, false)
    test(3, 2, 8, 0, false)
  })
})
