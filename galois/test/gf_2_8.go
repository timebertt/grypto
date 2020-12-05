package main

import (
  "fmt"

  "github.com/timebertt/grypto/galois"
)

func main() {
  f := galois.MustNewField(2, 8, "x^8 + x^4 + x^3 + x + 1")

  _a3 := f.MustNewElement(galois.Polynomial{1,1,0,0,0,1,0,1})
  _44 := f.MustNewElement(galois.Polynomial{0,0,1,0,0,0,1,0})
  fmt.Println(_a3.Add(_44))

  _c2 := f.MustNewElement(galois.Polynomial{0,1,0,0,0,0,1,1})
  fmt.Println(_c2.Inverse())
}
