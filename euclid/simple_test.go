package euclid_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/euclid"
)

var _ = Describe("#GreatestCommonDivisor", func() {
  It("should correctly calculate gcd", func() {
    test := func(a, b, expected int) {
      ExpectWithOffset(1, euclid.GreatestCommonDivisor(a, b)).To(Equal(expected), "gcd should be correct")
      ExpectWithOffset(1, euclid.GreatestCommonDivisor(b, a)).To(Equal(expected), "gcd (switched) should be correct")
    }

    test(0, 0, 0)
    test(1, 0, 1)
    test(2, 1, 1)
    test(3, 1, 1)
    test(3, 2, 1)
    test(15, 26, 1)
    test(30, 18, 6)
    test(-10, 20, 10)
    test(-20, -14, 2)
    test(235, 124, 1)
    test(3689, 3519, 17)
  })
})
