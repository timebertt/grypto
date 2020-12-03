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

func (f *Field) Zero() Element {
  return Element{Polynomial{0}, f}
}

func (f *Field) NewElement(p Polynomial) (Element, error) {
  e := Element{p, f}
  e.Normalize()

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
  for i := len(p) - 1; i >= 0; i-- {
    switch i {
    case 0:
      s += fmt.Sprintf(" + %d", p[i])
    case 1:
      s += fmt.Sprintf(" + %dx", p[i])
    default:
      s += fmt.Sprintf(" + %dx^%d", p[i], i)
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

type Element struct {
  Polynomial
  Field *Field
}

func (p Element) Copy() Element {
  return Element{p.Polynomial.Copy(), p.Field}
}

func (p Element) Add(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  sum := Element{p.Polynomial.Add(q.Polynomial), p.Field}
  return sum.Normalize()
}

func (p Element) Sub(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  qq := q.Copy()
  for i := range qq.Polynomial {
    qq.Polynomial[i] = -qq.Polynomial[i]
  }

  return p.Add(qq.Normalize())
}

func (p Element) Multiply(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  product := Element{p.Polynomial.Multiply(q.Polynomial), p.Field}
  return product.Normalize()
}

func (p Element) String() string {
  return p.Polynomial.String()
}

func (p Element) Divide(q Polynomial) (quotient Element, residue Element) {
  pp := p.Polynomial.Copy()
  pp.Normalize()
  qq := q.Copy()
  qq.Normalize()

  if qq.IsZero() {
    panic("grypto/galois: cannot divide by zero")
  }
  if pp.Degree() > qq.Degree() {
    panic("grypto/galois: cannot divide by a polynomial with a smaller degree")
  }

  s := make(Polynomial, pp.Degree()-qq.Degree()+1)
  r := make(Polynomial, pp.Degree()-qq.Degree()+1)

  j := pp.Degree()
  k := s.Degree()
  for i := qq.Degree(); i >= 0; i-- {
    r[i] += pp[j]
    _, _, rc := euclid.GreatestCommonDivisorExtended(int(r[i]), int(pp[j]))
    s[k] = int32(rc) * pp[j]
  }

  quotient = Element{s, p.Field}
  residue = Element{r, p.Field}
  return quotient.Normalize(), residue.Normalize()
}

func (p Element) Modulo(q Polynomial) Element {
  // _, r := p.Divide(q)
  // return r
  return p.Normalize()
}
