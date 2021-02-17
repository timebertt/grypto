package galois

import (
  "encoding/hex"
  "fmt"
  "math"
  "regexp"
  "strconv"
  "strings"
)

func MustNewField(p, n int32, modulus string) *Field {
  f, err := NewField(p, n, modulus)
  if err != nil {
    panic(fmt.Errorf("grypto/galois: not a valid field: %v", err))
  }
  return f
}

func NewField(p, n int32, modulus string) (*Field, error) {
  m, err := ParsePolynomial(modulus)
  if err != nil {
    return nil, fmt.Errorf("modulus not a valid polynomial: %w", err)
  }
  f := &Field{
    P: p,
    N: n,
  }
  f.Modulus = Element{Polynomial: m, Field: f}
  f.Modulus.Normalize()
  if f.Modulus.Degree() != n {
    return nil, fmt.Errorf("degree of modulus must be equal to field's n, but is %d", f.Modulus.Degree())
  }
  return f, nil
}

type Monomial struct {
  Coefficient int32
  Power       int32
}

func (m Monomial) ToPolynomial() Polynomial {
  p := make(Polynomial, m.Power+1)
  p[m.Power] = m.Coefficient
  return p
}

func (m Monomial) String() string {
  return m.ToPolynomial().String()
}

var monomialRegex = regexp.MustCompile(`^(-?[0-9]*)\*?(x)?(\^[0-9]*)?$`)

func MustParseMonomial(s string) Monomial {
  m, err := ParseMonomial(s)
  if err != nil {
    panic(fmt.Errorf("grypto/galois: not a valid monomial: %v", err))
  }
  return m
}

func ParseMonomial(s string) (Monomial, error) {
  s = strings.ReplaceAll(s, " ", "")
  matches := monomialRegex.FindStringSubmatch(s)
  if len(matches) == 0 || matches[0] == "" {
    return Monomial{}, fmt.Errorf("not a valid monomial: %q", s)
  }

  m := Monomial{}

  coefficient := matches[1]
  if coefficient == "" {
    m.Coefficient = 1
  } else if coefficient == "-" {
    m.Coefficient = -1
  } else {
    c, err := strconv.Atoi(coefficient)
    if err != nil {
      return Monomial{}, fmt.Errorf("not a valid coefficient: %s", coefficient)
    }
    if c > math.MaxInt32 {
      return Monomial{}, fmt.Errorf("coefficient exeeds max int32 value: %d", c)
    }
    m.Coefficient = int32(c)
  }

  x := matches[2]
  if x == "" {
    m.Power = 0
    return m, nil
  }

  power := strings.TrimPrefix(matches[3], "^")
  if power == "" {
    m.Power = 1
    return m, nil
  }

  p, err := strconv.Atoi(power)
  if err != nil {
    return Monomial{}, fmt.Errorf("not a valid power: %s", power)
  }
  if p > math.MaxInt32 {
    return Monomial{}, fmt.Errorf("power exeeds max int32 value: %d", p)
  }
  m.Power = int32(p)

  return m, nil
}

func MustParsePolynomial(s string) Polynomial {
  p, err := ParsePolynomial(s)
  if err != nil {
    panic(fmt.Errorf("grypto/galois: not a valid polynomial: %v", err))
  }
  return p
}

func ParsePolynomial(s string) (Polynomial, error) {
  var (
    coefficients = make(map[int32]int32, strings.Count(s, "+")+1)
    degree       = int32(0)
  )

  s = strings.ReplaceAll(s, " ", "")
  s = strings.ReplaceAll(s, "-", "+-")
  for i, part := range strings.Split(s, "+") {
    m, err := ParseMonomial(part)
    if err != nil {
      return nil, fmt.Errorf("failed to parse part %d as monomial: %w", i, err)
    }
    if _, ok := coefficients[m.Power]; ok {
      return nil, fmt.Errorf("multiple monomials with power %d specified", m.Power)
    }
    coefficients[m.Power] = m.Coefficient
    if m.Power > degree {
      degree = m.Power
    }
  }

  polynomial := make(Polynomial, degree+1)
  for p, c := range coefficients {
    polynomial[p] = c
  }
  polynomial.Normalize()

  return polynomial, nil
}

func MustParsePolynomialHex(s string) Polynomial {
  p, err := ParsePolynomialHex(s)
  if err != nil {
    panic(fmt.Errorf("grypto/galois: not a valid polynomial: %v", err))
  }
  return p
}

func ParsePolynomialHex(s string) (Polynomial, error) {
  s = strings.TrimPrefix(s, "0x")
  if len(s) == 0 {
    return Zero(), nil
  }
  if len(s)%2 != 0 {
    return nil, fmt.Errorf("can only parse polynomials with even number of hex digits")
  }

  bytes, err := hex.DecodeString(s)
  if err != nil {
    return nil, fmt.Errorf("error decoding hex string: %w", err)
  }

  p := make(Polynomial, len(s)*4)
  for i, b := range bytes {
    bit := byte(1)
    for j := 0; j < 8; j++ {
      if b&bit == bit {
        p[i*8+j] = int32(1)
      }
      bit <<= 1
    }
  }
  return p.Normalize(), nil
}

func (f *Field) ParseElement(s string) (Element, error) {
  p, err := ParsePolynomial(s)
  if err != nil {
    return Element{}, fmt.Errorf("not a valid polynomial: %w", err)
  }
  return f.NewElement(p)
}

func (f *Field) ParseElementHex(s string) (Element, error) {
  p, err := ParsePolynomialHex(s)
  if err != nil {
    return Element{}, fmt.Errorf("not a valid polynomial: %w", err)
  }
  return f.NewElement(p)
}

func (f *Field) MustParseElement(s string) Element {
  e, err := f.ParseElement(s)
  if err != nil {
    panic(fmt.Errorf("grypto/galois: not a valid polynomial: %v", err))
  }
  return e
}

func (f *Field) MustParseElementHex(s string) Element {
  e, err := f.ParseElementHex(s)
  if err != nil {
    panic(fmt.Errorf("grypto/galois: not a valid polynomial: %v", err))
  }
  return e
}

func (p Element) Normalize() Element {
  for i, c := range p.Polynomial {
    p.Polynomial[i] = (c + p.Field.P) % p.Field.P
  }
  p.Polynomial.Normalize()
  return p
}

func (p *Polynomial) Normalize() Polynomial {
  degree := len(*p) - 1
  for degree > 0 {
    if (*p)[degree] != 0 {
      break
    }
    degree--
  }

  *p = (*p)[:degree+1]
  return *p
}

func (p Polynomial) IsZero() bool {
  for _, c := range p {
    if c != 0 {
      return false
    }
  }
  return true
}

func (p Polynomial) IsOne() bool {
  if p[0] != 1 {
    return false
  }
  for _, c := range p[1:] {
    if c != 0 {
      return false
    }
  }
  return true
}
