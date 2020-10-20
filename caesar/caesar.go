package caesar

import (
  "crypto/cipher"
)

const (
  BlockSize = 1

  minUpper, maxUpper = 'A', 'Z'
  minLower, maxLower = 'a', 'z'

  modulus = 26
)

type caesar int

// NewCipher returns a new cipher.Block which implements the Caesar Cipher.
// It is a substituting block cipher operating on blocks of length 1 (single bytes).
// Latin characters are encrypted by replacing them by the `key`-th next character
// in the alphabet. The characters' case is kept and non-latin characters are not
// replaced.
// see https://en.wikipedia.org/wiki/Caesar_cipher
func NewCipher(key int) cipher.Block {
  // avoid overflow
  return caesar(key % modulus)
}

func (c caesar) BlockSize() int {
  return BlockSize
}

func (c caesar) Encrypt(dst, src []byte) {
  encryptBlock(int(c), dst, src)
}

func (c caesar) Decrypt(dst, src []byte) {
  encryptBlock(-int(c), dst, src)
}

func encryptBlock(key int, dst, src []byte) {
  if len(src) < BlockSize {
    panic("grypto/caesar: input not full block")
  }
  if len(dst) < BlockSize {
    panic("grypto/caesar: output not full block")
  }

  // modulo operation does not return a positive residue for a negative number
  if key < 0 {
    key += modulus
  }

  // convert to rune for easy comparing
  in := rune(src[0])

  switch {
  case key == 0:
    fallthrough
  default:
    dst[0] = src[0]
  case in >= minUpper && in <= maxUpper:
    dst[0] = byte(substitute(key, in, minUpper))
  case in >= minLower && in <= maxLower:
    dst[0] = byte(substitute(key, in, minLower))
  }

  return
}

func substitute(key int, in, min rune) rune {
  return min + rune((int(in-min)+key)%modulus)
}
