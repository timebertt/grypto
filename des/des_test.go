package des

import (
  "fmt"
  "math/rand"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("#permuteInitial", func() {
  It("should permute block correctly", func() {
    for i := uint(0); i < 64; i++ {
      // use single bit at position i as input
      bit := uint64(1) << i
      // look up where bit i should get shifted to in inverse permutation
      // (rather than looping over the original permutation and finding i)
      expected := uint64(1) << finalPermutation[63-i]

      permuted := permuteInitial(bit)
      Expect(permuted).To(Equal(expected), fmt.Sprintf("i=%d", i))
    }
  })
  It("should reverse final permutation", func() {
    block := rand.Uint64()
    Expect(permuteInitial(permuteFinal(block))).To(Equal(block))
  })
})

var _ = Describe("#permuteFinal", func() {
  It("should permute block correctly", func() {
    for i := uint(0); i < 64; i++ {
      // use single bit at position i as input
      bit := uint64(1) << i
      // look up where bit i should get shifted to in inverse permutation
      // (rather than looping over the original permutation and finding i)
      expected := uint64(1) << initialPermutation[63-i]

      permuted := permuteFinal(bit)
      Expect(permuted).To(Equal(expected), fmt.Sprintf("i=%d", i))
    }
  })
  It("should reverse initial permutation", func() {
    block := rand.Uint64()
    Expect(permuteFinal(permuteInitial(block))).To(Equal(block))
  })
})

var _ = Describe("#validateKey", func() {
  It("should return error, if key does not have a size of 8", func() {
    Expect(validateKey([]byte{0, 1, 2, 3, 4, 5, 6})).To(MatchError(ContainSubstring("invalid key size")))
    Expect(validateKey([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})).To(MatchError(ContainSubstring("invalid key size")))
  })
  It("should return error, if one bytes has even parity", func() {
    Expect(validateKey([]byte{
      0b11011100, 0b01110110, 0b11110100, 0b01011000,
      0b11010110, 0b01111010, 0b10110101, 0b00001001,
    })).To(MatchError(ContainSubstring("byte 7 has even parity")))
  })
  It("should succeed, if all bytes have odd parity", func() {
    Expect(validateKey([]byte{
      0b11011100, 0b01110110, 0b11110100, 0b01011000,
      0b11010110, 0b01111010, 0b10110101, 0b00001000,
    })).To(Succeed())
  })
})
