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

func NewCipher() cipher.Block {
  return &des{}
}

func (d *des) BlockSize() int {
  return BlockSize
}

func (d *des) Encrypt(dst, src []byte) {

}

func (d *des) Decrypt(dst, src []byte) {
  panic("implement me")
}

func permute(p [64]byte, src uint64) (dst uint64) {
  // start with rightmost bit (position 63 in src)
  bit := uint64(1)

  for i := 0; i < 64; i++ {
    dst |= (src << p[63-i])&bit

    // one step left
    bit <<= i
  }
  return
}

func feistel(rounds int, dst, src []byte) {

}
