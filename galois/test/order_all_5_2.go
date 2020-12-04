package main

import (
  "fmt"

  "github.com/timebertt/grypto/galois"
)

func main() {
  f := galois.MustNewField(5, 2, "x^2 + 2")

  for i := int32(0); i < 5; i++ {
    for j := int32(0); j < 5; j++ {
      e := f.MustNewElement(galois.Polynomial{i, j})
      fmt.Printf("order(%s) = %d\n", e, e.Order())
    }
  }
}
