package galois_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/galois"
)

var _ = Describe("ParseMonomial", func() {
  It("should return error for invalid monomials", func() {
    test := func(s, expectedError string) {
      m, err := galois.ParseMonomial(s)
      ExpectWithOffset(1, m).To(BeZero())
      ExpectWithOffset(1, err).To(MatchError(ContainSubstring(expectedError)))
    }

    test("", "not a valid monomial")
    test("foo", "not a valid monomial")
    test("2x^2x", "not a valid monomial")
  })
  It("should correctly parse different monomials", func() {
    test := func(s string, coefficient, power int32) {
      m, err := galois.ParseMonomial(s)
      ExpectWithOffset(1, err).NotTo(HaveOccurred())
      ExpectWithOffset(1, m).To(Equal(galois.Monomial{
        Coefficient: coefficient,
        Power:       power,
      }))
    }

    test("0", 0, 0)
    test("1", 1, 0)
    test("123", 123, 0)
    test(" 123 ", 123, 0)
    test("x^0", 1, 0)
    test("x ^ 0", 1, 0)
    test("x", 1, 1)
    test("1x", 1, 1)
    test("123x", 123, 1)
    test("123*x", 123, 1)
    test("123 * x", 123, 1)
    test("x^1", 1, 1)
    test("x^123", 1, 123)
    test("x ^ 123", 1, 123)
    test("456x^123", 456, 123)
    test("456*x^123", 456, 123)
    test("456 * x ^ 123", 456, 123)
  })
})

var _ = Describe("ParsePolynomial", func() {
  It("should return error for invalid polynomials", func() {
    test := func(s, expectedError string) {
      p, err := galois.ParsePolynomial(s)
      ExpectWithOffset(1, p).To(BeNil())
      ExpectWithOffset(1, err).To(MatchError(ContainSubstring(expectedError)))
    }

    test("", "not a valid monomial")
    test("foo", "not a valid monomial")
    test("1+1", "multiple monomials with power 0")
    test("x+x", "multiple monomials with power 1")
    test("x^4+x^4", "multiple monomials with power 4")
  })
  It("should correctly parse different polynomials", func() {
    test := func(s string, expected galois.Polynomial) {
      p, err := galois.ParsePolynomial(s)
      ExpectWithOffset(1, err).NotTo(HaveOccurred())
      ExpectWithOffset(1, p).To(Equal(expected), "should match "+expected.String())
    }

    test("0", []int32{0})
    test("1", []int32{1})
    test("123", []int32{123})
    test(" 123 ", []int32{123})
    test("x^0", []int32{1})
    test("x ^ 0", []int32{1})
    test("x", []int32{0, 1})
    test("x+1", []int32{1, 1})
    test("1x", []int32{0, 1})
    test("1*x", []int32{0, 1})
    test("1*x + 123", []int32{123, 1})
    test("1 * x", []int32{0, 1})
    test("123x", []int32{0, 123})
    test("x^1", []int32{0, 1})
    test("x ^ 1", []int32{0, 1})
    test("x^2", []int32{0, 0, 1})
    test("456x^4", []int32{0, 0, 0, 0, 456})
    test("0x^4", []int32{0})
    test("456 * x ^ 4", []int32{0, 0, 0, 0, 456})
    test("456 * x ^ 4 + 4x^2", []int32{0, 0, 4, 0, 456})
    test("4x^2 + 456 * x ^ 4", []int32{0, 0, 4, 0, 456})
  })
})

var _ = Describe("NewField", func() {
  It("should return error for invalid fields", func() {
    test := func(p, n int32, m string, expectedError string) {
      f, err := galois.NewField(p, n, m)
      ExpectWithOffset(1, f).To(BeZero())
      ExpectWithOffset(1, err).To(MatchError(ContainSubstring(expectedError)))
    }

    test(2, 8, "foo", "not a valid polynomial")
    test(2, 8, "4x^3+x^2+3", "degree of modulus")
    test(2, 2, "4x^3+x^2+3", "degree of modulus")
  })
  It("should correctly construct new field", func() {
    test := func(p, n int32, m string, expectedModulus string) {
      f, err := galois.NewField(p, n, m)
      ExpectWithOffset(1, err).NotTo(HaveOccurred())
      ExpectWithOffset(1, f.P).To(Equal(p))
      ExpectWithOffset(1, f.N).To(Equal(n))

      expected := galois.MustParsePolynomial(expectedModulus)
      ExpectWithOffset(1, f.Modulus).To(Equal(expected), "should match "+expectedModulus)
    }

    test(5, 3, "5x^3+6x^2+3", "1x^2+3")
  })
})

var _ = Describe("ParseElement", func() {
  var (
    field *galois.Field
  )

  BeforeEach(func() {
    field = galois.MustNewField(5, 2, "3x^2+4x+1")
  })

  It("should return error for invalid elements", func() {
    test := func(s, expectedError string) {
      e, err := field.ParseElement(s)
      ExpectWithOffset(1, e).To(BeZero())
      ExpectWithOffset(1, err).To(MatchError(ContainSubstring(expectedError)))
    }

    test("", "not a valid monomial")
    test("foo", "not a valid monomial")
    test("1+1", "multiple monomials with power 0")
    test("x+x", "multiple monomials with power 1")
    test("x^4+x^4", "multiple monomials with power 4")
    test("x^4+x^2", "degree of polynomial is larger")
  })
  It("should correctly parse different elements", func() {
    test := func(s string, expected galois.Polynomial) {
      e, err := field.ParseElement(s)
      ExpectWithOffset(1, err).NotTo(HaveOccurred())
      ExpectWithOffset(1, e.Field).To(BeIdenticalTo(field))
      ExpectWithOffset(1, e.Polynomial).To(Equal(expected), "should match "+expected.String())
    }

    test("0", []int32{0})
    test("1", []int32{1})
    test("8", []int32{3})
    test("x", []int32{0, 1})
    test("x+12", []int32{2, 1})
    test("x+10", []int32{0, 1})
    test("5x^2+x+10", []int32{0, 1})
  })
})
