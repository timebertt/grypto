package des

import (
  "crypto/cipher"
)

const (
  BlockSize = 8 // 64 bit
)

type des struct {
  key uint64
}

func NewCipher(key uint64) cipher.Block {
  return &des{key}
}

func (d *des) BlockSize() int {
  return BlockSize
}

func (d *des) Encrypt(dst, src []byte) {

}

func (d *des) Decrypt(dst, src []byte) {
  panic("implement me")
}

func permuteInitial(block uint64) uint64 {
  // use finalPermutation to look up where each bit should be shifted to
  // instead of looping through initialPermutation to find i
  return permute(finalPermutation, block)
}

func permuteFinal(block uint64) uint64 {
  // use initialPermutation to look up where each bit should be shifted to
  // instead of looping through finalPermutation to find i
  return permute(initialPermutation, block)
}

func permute(p [64]byte, src uint64) (dst uint64) {
  // start with rightmost bit (position 63 in src)
  bit := uint64(1)

  for i := 0; i < 64; i++ {
    shiftBy := int(p[63-i]) - i
    if shiftBy > 0 {
      dst |= (src & bit) << shiftBy
    } else {
      dst |= (src & bit) >> -shiftBy
    }

    // one step left
    bit <<= 1
  }
  return
}

func feistel(rounds int, dst, src []byte) {

}
