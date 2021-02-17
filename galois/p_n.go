package galois

import (
  "encoding/hex"
  "fmt"
  "math"
  "reflect"
  "strings"

  "github.com/timebertt/grypto/euclid"
)

var AESField = MustNewField(2, 8, "x^8 + x^4 + x^3 + x + 1")

type Field struct {
  P, N    int32
  Modulus Element
}

func Zero() Polynomial {
  return Polynomial{0}
}

func One() Polynomial {
  return Polynomial{1}
}

func (f *Field) Zero() Element {
  return f.uncheckedElement(Zero())
}

func (f *Field) One() Element {
  return f.uncheckedElement(One())
}

func (f *Field) uncheckedElement(p Polynomial) Element {
  e := Element{p, f}
  return e.Normalize()
}

func (f *Field) MustNewElement(p Polynomial) Element {
  e, err := f.NewElement(p)
  if err != nil {
    panic(fmt.Errorf("grypto/galois: not a valid element: %v", err))
  }
  return e
}

func (f *Field) Order() int32 {
  return int32(math.Pow(float64(f.P), float64(f.N))) - 1
}

func (f *Field) NewElement(p Polynomial) (Element, error) {
  e := f.uncheckedElement(p)
  if p.Degree() > f.N {
    return Element{}, fmt.Errorf("degree of polynomial is larger than field's N")
  }
  return e, nil
}

type Polynomial []int32

func (p Polynomial) Copy() Polynomial {
  return append(p[:0:0], p...)
}

func (p Polynomial) String() string {
  s := ""
  printCoefficient := func(i int32, leaveOutOne bool) string {
    switch {
    case leaveOutOne && i == 1:
      return "+ "
    case leaveOutOne && i == -1:
      return "- "
    case i < 0:
      return fmt.Sprintf("- %d", -i)
    }
    return fmt.Sprintf("+ %d", i)
  }

  for i := len(p) - 1; i >= 0; i-- {
    c := p[i]

    if c == 0 {
      continue
    }
    switch i {
    case 0:
      s += fmt.Sprintf(" %s", printCoefficient(c, false))
    case 1:
      s += fmt.Sprintf(" %sx", printCoefficient(c, true))
    default:
      s += fmt.Sprintf(" %sx^%d", printCoefficient(c, true), i)
    }

    if i == len(p)-1 {
      s = strings.Replace(s, " + ", "", 1)
      s = strings.Replace(s, " - ", "-", 1)
    }
  }
  if len(s) == 0 {
    return "0"
  }
  return s
}

func (p Polynomial) Degree() int32 {
  for i := int32(len(p)) - 1; i >= 0; i-- {
    if p[i] != 0 {
      return i
    }
  }
  return -1
}

func (p Polynomial) Lead() int32 {
  for i := p.Degree(); i >= 0; i-- {
    if p[i] != 0 {
      return p[i]
    }
  }
  return 0
}

func (p Polynomial) Add(q Polynomial) Polynomial {
  maxDeg := p.Degree()
  if q.Degree() > maxDeg {
    maxDeg = q.Degree()
  }
  // result polynomial: deg r = max (deg p, deg q)
  r := make(Polynomial, maxDeg+1)

  for i := 0; i <= int(maxDeg); i++ {
    if i < len(p) {
      r[i] += p[i]
    }
    if i < len(q) {
      r[i] += q[i]
    }
  }
  return r
}

func (p Polynomial) Sub(q Polynomial) Polynomial {
  q = q.Copy()
  for i := range q {
    q[i] = -q[i]
  }
  d := p.Add(q.Normalize())
  return d.Normalize()
}

func (p Polynomial) Multiply(q Polynomial) Polynomial {
  if p.IsZero() || q.IsZero() {
    return Zero()
  }

  // result polynomial: deg r = deg p + deg q
  r := make(Polynomial, p.Degree()+q.Degree()+1)
  for i, pi := range p {
    for j, qj := range q {
      r[i+j] = r[i+j] + pi*qj
    }
  }
  r.Normalize()
  return r
}

type Element struct {
  Polynomial
  Field *Field
}

func (p Element) String() string {
  return p.Polynomial.String()
}

func (p Element) HexString() string {
  if p.Field.P != 2 || p.Field.N != 8 {
    panic("grypto/galois: HexString is only supported for elements of GF(2^8)")
  }

  var b byte
  for i, c := range p.Polynomial {
    b |= byte(c) << i
  }
  return "0x" + hex.EncodeToString([]byte{b})
}

func (p Element) Copy() Element {
  return p.Field.uncheckedElement(p.Polynomial.Copy())
}

func (p Element) Add(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  return p.Field.uncheckedElement(p.Polynomial.Add(q.Polynomial))
}

func (p Element) Sub(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  q = q.Copy()
  for i := range q.Polynomial {
    q.Polynomial[i] = -q.Polynomial[i]
  }

  return p.Add(q.Normalize()).Normalize()
}

func (p Element) Multiply(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  return p.Field.uncheckedElement(p.Polynomial.Multiply(q.Polynomial)).Modulo(p.Field.Modulus)
}

func (p Element) Divide(q Element) (quotient, residue Element) {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }
  if q.IsOne() {
    return q.Normalize(), p.Field.Zero()
  }

  p = p.Copy()
  p.Normalize()
  q = q.Copy()
  q.Normalize()

  if q.IsZero() {
    panic("grypto/galois: cannot divide by zero")
  }
  if p.Degree() < q.Degree() {
    return p.Field.Zero(), p
  }

  quotient = Element{Polynomial: make(Polynomial, p.Degree()-q.Degree()+1), Field: p.Field}
  residue = p.Copy()

  i := p.Degree()
  j := q.Degree()
  for i >= j {
    _, _, qLeadInv := euclid.GreatestCommonDivisorExtended(int(p.Field.P), int(q.Lead()))
    qLeadInvNormalized := (int32(qLeadInv) + p.Field.P) % p.Field.P

    quotient.Polynomial[i-j] = residue.Lead() * qLeadInvNormalized
    quotientMonomial := Monomial{Coefficient: quotient.Polynomial[i-j], Power: i - j}

    residue = residue.Sub(p.Field.uncheckedElement(q.Polynomial.Multiply(quotientMonomial.ToPolynomial())))
    i = residue.Degree()
  }

  return quotient.Normalize(), residue.Normalize()
}

func (p Element) Modulo(q Element) Element {
  _, r := p.Divide(q)
  return r
}

func (p Element) Pow(e int32) Element {
  if p.IsZero() {
    return p
  }

  e = (e + p.Field.Order()) % p.Field.Order()

  exp := p.Field.One()
  for i := int32(1); i <= e; i++ {
    exp = exp.Multiply(p)
  }
  return exp
}

func (p Element) Order() int32 {
  if p.IsZero() {
    return -1
  }

  exp := p.Field.One()
  for i := int32(1); i <= p.Field.Order(); i++ {
    exp = exp.Multiply(p)
    if exp.IsOne() {
      return i
    }
  }
  return -1
}

func (p Element) Inverse() Element {
  _, _, inv := GreatestCommonDivisorExtended(p.Field.Modulus, p)
  return inv
}

func GreatestCommonDivisorExtended(a, b Element) (gcd, x, y Element) {
  if a.Field != b.Field && reflect.DeepEqual(a.Field, b.Field) {
    panic("grypto/galois: a and b are not element of the same field")
  }

  if a.Degree() < b.Degree() {
    panic("grypto/galois: b's degree must not be greater than a's degree")
  }
  return extendedEuclid(a, b)
}

func extendedEuclid(a, b Element) (gcd, x, y Element) {
  var (
    x0, x1 = a.Field.One(), a.Field.Zero()
    y0, y1 = a.Field.Zero(), a.Field.One()
    q, r   Element
  )

  for !b.IsZero() {
    q, r = a.Divide(b)
    a, b = b, r
    x0, x1 = x1, x0.Sub(q.Multiply(x1))
    y0, y1 = y1, y0.Sub(q.Multiply(y1))
  }
  return a.Normalize(), x0, y0
}
