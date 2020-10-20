package euclid_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/euclid"
)

var _ = Describe("#GreatestCommonDivisorExtended", func() {
  It("should correctly calculate gcd and linear combination", func() {
    test := func(a, b, expectedGCD, expectedX, expectedY int) {
      gcd, x, y := euclid.GreatestCommonDivisorExtended(a, b)
      ExpectWithOffset(1, gcd).To(Equal(expectedGCD), "gcd should be correct")
      ExpectWithOffset(1, x).To(Equal(expectedX), "x should be correct")
      ExpectWithOffset(1, y).To(Equal(expectedY), "y should be correct")
    }

    test(0, 0, 0, 1, 0)
    test(1, 0, 1, 1, 0)
    test(2, 1, 1, 0, 1)
    test(3, 1, 1, 0, 1)
    test(3, 2, 1, 1, -1)
    test(15, 26, 1, 7, -4)
    test(30, 18, 6, -1, 2)
    test(235, 124, 1, 19, -36)
    test(3689, 3519, 17, -62, 65)
  })
})
