package galois

import (
  "fmt"
  "reflect"

  "github.com/timebertt/grypto/euclid"
)

type Field struct {
  P, N    int32
  Modulus Polynomial
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
  }
  if len(s) == 0 {
    return "0"
  }
  return s[3:]
}

func (p Polynomial) Degree() int32 {
  return int32(len(p) - 1)
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

func (p Polynomial) Multiply(q Polynomial) Polynomial {
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

func (p Polynomial) Divide(q Polynomial) (quotient Polynomial, residue Polynomial) {
  if q.IsOne() {
    return q.Normalize(), Zero()
  }

  p = p.Copy()
  p.Normalize()
  q = q.Copy()
  q.Normalize()

  if q.IsZero() {
    panic("grypto/galois: cannot divide by zero")
  }
  if p.Degree() > q.Degree() {
    panic("grypto/galois: cannot divide by a polynomial with a smaller degree")
  }

  quotient = make(Polynomial, p.Degree()-q.Degree()+1)
  residue = p.Copy()

  var j int32
  k := quotient.Degree()
  for i := q.Degree(); i >= 0; i-- {
    j = residue.Degree()
    _, _, rc := euclid.GreatestCommonDivisorExtended(int(residue[i]), int(p[j]))
    quotient[k] = int32(rc) * p[j]
  }

  return quotient.Normalize(), residue.Normalize()
}

func (p Polynomial) Modulo(q Polynomial) Polynomial {
  // _, r := p.Divide(q)
  // return r
  return p.Normalize()
}

type Element struct {
  Polynomial
  Field *Field
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

  return p.Add(q.Normalize())
}

func (p Element) Multiply(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  return p.Field.uncheckedElement(p.Polynomial.Multiply(q.Polynomial))
}

func (p Element) String() string {
  return p.Polynomial.String()
}
