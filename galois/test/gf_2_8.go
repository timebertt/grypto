package main

import (
  "fmt"

  "github.com/timebertt/grypto/galois"
)

func main() {
  f := galois.AESField

  a := f.MustParseElementHex("a3")
  fmt.Println(a)
  fmt.Println(a.HexString())
  b := f.MustParseElementHex("44")
  fmt.Println(b)
  fmt.Println(b.HexString())
  fmt.Println(a.Add(b))
  fmt.Println(a.Add(b).HexString())

  a = f.MustParseElementHex("12")
  fmt.Println(a)
  fmt.Println(a.HexString())
  b = f.MustParseElementHex("f9")
  fmt.Println(b)
  fmt.Println(b.HexString())
  fmt.Println(a.Add(b))
  fmt.Println(a.Add(b).HexString())

  c := f.MustParseElementHex("c2")
  fmt.Println(c)
  fmt.Println(c.Inverse())
  fmt.Println(c.Inverse().HexString())
  fmt.Println(c.Inverse().Inverse().HexString())
}
