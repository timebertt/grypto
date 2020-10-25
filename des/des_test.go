package des

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("#initialPermutation", func() {
  It("should permute block correctly", func() {
    for i := uint(0); i < 64; i++ {
      // use single bit at position i as input
      bit := uint64(1) << i
      // look up where bit i should get shifted to in inverse permutation
      // (rather than looping over the original permutation and finding i)
      expected := uint64(1) << finalPermutation[63-i]

      permuted := permute(initialPermutation, bit)
      Expect(permuted).To(Equal(expected))
    }
  })
})
