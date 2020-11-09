package des

import (
  "crypto/cipher"
  "encoding/binary"
  "fmt"
)

const (
  BlockSize = 8 // 64 bit
)

type KeySizeError int

func (k KeySizeError) Error() string {
  return fmt.Sprintf("grypto/des: invalid key size: %d", int(k))
}

type des struct {
  key []byte
}

func NewCipher(key []byte) (cipher.Block, error) {
  if err := validateKey(key); err != nil {
    return nil, err
  }
  return &des{key}, nil
}

func validateKey(key []byte) error {
  if len(key) != 8 {
    return KeySizeError(len(key))
  }

  for i, b := range key {
    bit := uint8(1) << 7
    sum := uint8(0)
    for bit > 0 {
      if bit&b == bit {
        sum++
      }
      bit >>= 1
    }

    if sum%2 == 0 {
      return fmt.Errorf("grypto/des: invalid key: byte %d has even parity: %b", i, b)
    }
  }

  return nil
}

func (d *des) BlockSize() int {
  return BlockSize
}

func (d *des) Encrypt(dst, src []byte) {
  encrypt(d.key, dst, src)
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

func encrypt(key, dst, src []byte) {
  if len(src) < BlockSize {
    panic("grypto/des: input not full block")
  }
  if len(dst) < BlockSize {
    panic("grypto/des: output not full block")
  }

  feistel(16, key, dst, src)
}

func feistel(rounds int, key, dst, src []byte) {
  b := binary.BigEndian.Uint64(src)
  b = permuteInitial(b)
  left, right := uint32(b>>32), uint32(b)

  for i := 0; i < rounds; i++ {
    left, right = right, left^f(right, key)
  }

  b = uint64(left)<<32 | uint64(right)
  binary.BigEndian.PutUint64(dst, permuteFinal(b))
}

func f(in uint32, key []byte) uint32 {

}

func expand() {

}
