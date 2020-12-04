package main

import (
  "fmt"

  "github.com/timebertt/grypto/galois"
)

func main() {
  f := galois.MustNewField(5, 2, "x^2 + 2")

  var all []galois.Element

  for i := int32(0); i < 5; i++ {
    for j := int32(0); j < 5; j++ {
      all = append(all, f.MustNewElement(galois.Polynomial{i, j}))
    }
  }

  for i := 0; i < len(all); i++ {
    for j := i; j < len(all); j++ {
      e1, e2 := all[i], all[j]
      product := e1.Multiply(e2)
      fmt.Printf("(%s) * (%s) = %s\n", e1.String(), e2.String(), product.String())
    }
  }
}
