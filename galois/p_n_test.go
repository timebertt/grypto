package galois_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/galois"
)

var _ = Describe("Polynomial", func() {
  Describe("#Add", func() {
    It("should correctly calculate sum", func() {
      test := func(p, q, expected string) {
        pp := galois.MustParsePolynomial(p)
        qq := galois.MustParsePolynomial(q)
        e := galois.MustParsePolynomial(expected)
        ExpectWithOffset(1, pp.Add(qq)).To(Equal(e), "should be equal to "+e.String())
      }

      test("0", "0", "0")
      test("0", "1", "1")
      test("0", "3x+1", "3x+1")
      test("1", "3x+1", "3x+2")
      test("2x+1", "3x+1", "5x+2")
      test("2x^2+3", "4x+5", "2x^2 + 4x + 8")
    })
  })

  Describe("#Multiply", func() {
    It("should correctly calculate product", func() {
      test := func(p, q, expected string) {
        pp := galois.MustParsePolynomial(p)
        qq := galois.MustParsePolynomial(q)
        e := galois.MustParsePolynomial(expected)
        ExpectWithOffset(1, pp.Multiply(qq)).To(Equal(e), "should be equal to "+e.String())
      }

      test("0", "0", "0")
      test("0", "1", "0")
      test("0", "3x+1", "0")
      test("1", "3x+1", "3x+1")
      test("2x+1", "3x+1", "6x^2 + 5x + 1")
      test("2x^2+3", "4x+5", "8x^3 + 10x^2 + 12x + 15")
    })
  })

})
var _ = Describe("Element", func() {
  var (
    field *galois.Field
  )

  BeforeEach(func() {
    field = galois.MustNewField(5,2, "3x^2 + 4x + 1")
  })

  Describe("#Add", func() {
    It("should correctly calculate sum", func() {
      test := func(p, q, expected string) {
        pp := field.MustParseElement(p)
        qq := field.MustParseElement(q)
        e := field.MustParseElement(expected)
        ExpectWithOffset(1, pp.Add(qq)).To(Equal(e), "should be equal to "+e.String())
      }

      test("0", "0", "0")
      test("0", "1", "1")
      test("4", "2", "1")
      test("0", "3x+1", "3x+1")
      test("1", "3x+1", "3x+2")
      test("2x+1", "3x+1", "2")
      test("2x^2+3", "4x+5", "2x^2 + 4x + 8")
      test("2x^2+3", "4x^2+5", "1x^2 + 0x + 3")
    })
  })

  Describe("#Multiply", func() {
    It("should correctly calculate product", func() {
      test := func(p, q, expected string) {
        pp := field.MustParseElement(p)
        qq := field.MustParseElement(q)
        e := field.MustParseElement(expected)
        ExpectWithOffset(1, pp.Multiply(qq)).To(Equal(e), "should be equal to "+e.String())
      }

      test("0", "0", "0")
      test("0", "1", "0")
      test("0", "3x+1", "0")
      test("1", "3x+1", "3x+1")
      test("2x+1", "3x+1", "1x^2 + 0x + 1")
      // test("2x^2+3", "4x+5", "calculate me")
    })
  })
})
