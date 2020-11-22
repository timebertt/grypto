package modular_test

import (
  "math"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/modular"
)

var _ = Describe("SubgroupOf", func() {
  It("should panic on invalid inputs", func() {
    test := func(b, m int32) {
      ExpectWithOffset(1, func() {
        modular.SubgroupOf(b, m)
      }).To(Panic())
    }

    test(-1, 2)
    test(1, -1)
    test(0, 2)
    test(1, 0)
    test(1, math.MaxInt32)
  })

  It("should correctly calculate subgroup", func() {
    test := func(b, m int32, expected []int32) {
      s := modular.SubgroupOf(b, m)
      ExpectWithOffset(1, s).To(Equal(expected))
    }

    test(1, 13, []int32{1})
    test(3, 13, []int32{1, 3, 9})
    test(4, 13, []int32{1, 4, 3, 12, 9, 10})
    test(10, 13, []int32{1, 10, 9, 12, 3, 4})
    test(3, 7, []int32{1, 3, 2, 6, 4, 5})
    test(2, 7, []int32{1, 2, 4})
    test(2, 8, []int32{1, 2, 4, 0})
  })
})
