package galois_test

import (
  "fmt"

  "github.com/timebertt/grypto/galois"
)

func ExampleElement_Multiply() {
  f := galois.MustNewField(5,2,"3x^2 + 4x + 1")
  e1 := f.MustParseElement("2x+1")
  e2 := f.MustParseElement("3x+1")

  fmt.Println(e1.String())
  fmt.Println(e2.String())
  fmt.Println(e1.Multiply(e2).String())
}
