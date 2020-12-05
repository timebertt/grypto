package aes

// Modulus is the polynomial which is used to construct GF(2⁸) in AES.
const Modulus = 1<<8 | 1<<4 | 1<<3 | 1<<1 | 1<<0 // x⁸ + x⁴ + x³ + x + 1

// Element is an element (polynomial) in the AES field.
type Element byte

func Add(a, b Element) Element {
  panic("implement me")
}

func Sub(a, b Element) Element {
  return Add(a, b)
}

func Mul(a, b Element) Element {
  panic("implement me")
}

func Inv(a, b Element) Element {
  panic("implement me")
}
