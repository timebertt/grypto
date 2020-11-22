package modular_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/modular"
)

var _ = Describe("OrderOf", func() {
  It("should panic on invalid inputs", func() {
    test := func(b, m int32) {
      ExpectWithOffset(1, func() {
        modular.OrderOf(b, m)
      }).To(Panic())
    }

    test(-1, 2)
    test(1, -1)
    test(0, 2)
    test(1, 0)
  })

  It("should correctly calculate order", func() {
    test := func(b, m, expected int32, expectedInf bool) {
      o, inf := modular.OrderOf(b, m)
      ExpectWithOffset(1, inf).To(Equal(expectedInf))
      if !expectedInf {
        ExpectWithOffset(1, o).To(Equal(expected))
      }
    }

    test(1, 13, 1, false)
    test(3, 13, 3, false)
    test(4, 13, 6, false)
    test(10, 13, 6, false)
    test(3, 7, 6, false)
    test(2, 7, 3, false)
    test(3, 9, 0, true)
    test(2, 8, 0, true)
  })
})
