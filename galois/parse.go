package galois

import (
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
    P:       p,
    N:       n,
    Modulus: m,
  }
  f.NormalizePolynomial(&f.Modulus)
  if m.Degree() != n {
    return nil, fmt.Errorf("degree of modulus must be equal to field's n, but is %d", m.Degree())
  }
  return f, nil
}

type Monomial struct {
  Coefficient int32
  Power       int32
}

var monomialRegex = regexp.MustCompile(`^([0-9]*)\*?(x)?(\^[0-9]*)?$`)

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

  for i, part := range strings.Split(strings.ReplaceAll(s, " ", ""), "+") {
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

func (f *Field) ParseElement(s string) (Element, error) {
  p, err := ParsePolynomial(s)
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

func (f *Field) NormalizePolynomial(p *Polynomial) {
  for i, c := range *p {
    (*p)[i] = (c + f.P) % f.P
  }
  p.Normalize()
}

func (p *Element) Normalize() Element {
  for i, c := range p.Polynomial {
    p.Polynomial[i] = (c + p.Field.P) % p.Field.P
  }
  p.Polynomial.Normalize()
  return *p
}

func (p *Polynomial) Normalize() Polynomial {
  if p.Degree() == 0 {
    return *p
  }

  degree := p.Degree()
  for degree > 0 {
    if (*p)[degree] > 0 {
      break
    }
    degree--
  }

  *p = (*p)[:degree+1]
  return *p
}

func (p Polynomial) IsZero() bool {
  for _, c := range p {
    if c > 0 {
      return true
    }
  }
  return false
}
