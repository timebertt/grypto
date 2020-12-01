package galois

import (
  "fmt"
  "reflect"
)

type Field struct {
  P, N    int32
  Modulus Polynomial
}

func (f *Field) Null() Element {
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

func (p Element) Add(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  return Element{p.Polynomial.Add(q.Polynomial), p.Field}.Modulo(q.Field.Modulus)
}

func (p Element) Multiply(q Element) Element {
  if p.Field != q.Field && reflect.DeepEqual(p.Field, q.Field) {
    panic("grypto/galois: p and q are not element of the same field")
  }

  return Element{p.Polynomial.Multiply(q.Polynomial), p.Field}.Modulo(q.Field.Modulus)
}

func (p Element) String() string {
  return p.Polynomial.String()
}

func (p Element) Divide(q Polynomial) (quotient Element, residue Element) {
  return p, p.Field.Null()
}

func (p Element) Modulo(q Polynomial) Element {
  // _, r := p.Divide(q)
  // return r
  return p.Normalize()
}
