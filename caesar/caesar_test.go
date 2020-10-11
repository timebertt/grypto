package caesar_test

import (
  "crypto/cipher"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/timebertt/grypto/caesar"
)

var _ = Describe("Caesar Cipher", func() {
  var (
    c        cipher.Block
    src, dst []byte
  )

  BeforeEach(func() {
    src = make([]byte, 1)
    dst = make([]byte, 1)
  })

  Context("key > 0", func() {
    BeforeEach(func() {
      c = caesar.NewCipher(1)
    })

    It("should encrypt upper case letter", func() {
      in := byte('A')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('B')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should encrypt lower case letter", func() {
      in := byte('a')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('b')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should not encrypt non-latin letter", func() {
      in := byte('!')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(in))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })
  })

  Context("key > 26", func() {
    BeforeEach(func() {
      c = caesar.NewCipher(27)
    })

    It("should encrypt upper case letter", func() {
      in := byte('A')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('B')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should encrypt lower case letter", func() {
      in := byte('a')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('b')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should not encrypt non-latin letter", func() {
      in := byte('!')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(in))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })
  })

  Context("key < 0", func() {
    BeforeEach(func() {
      c = caesar.NewCipher(-2)
    })

    It("should encrypt upper case letter", func() {
      in := byte('A')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('Y')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should encrypt lower case letter", func() {
      in := byte('a')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('y')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should not encrypt non-latin letter", func() {
      in := byte('!')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(in))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })
  })

  Context("key < 26", func() {
    BeforeEach(func() {
      c = caesar.NewCipher(-28)
    })

    It("should encrypt upper case letter", func() {
      in := byte('A')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('Y')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should encrypt lower case letter", func() {
      in := byte('a')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(byte('y')))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should not encrypt non-latin letter", func() {
      in := byte('!')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(in))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })
  })

  Context("key == 0", func() {
    BeforeEach(func() {
      c = caesar.NewCipher(0)
    })

    It("should not encrypt upper case letter", func() {
      in := byte('A')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(in))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should not encrypt lower case letter", func() {
      in := byte('a')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(in))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })

    It("should not encrypt non-latin letter", func() {
      in := byte('!')
      src[0] = in
      c.Encrypt(dst, src)
      Expect(dst).To(ConsistOf(in))
      c.Decrypt(src, dst)
      Expect(src).To(ConsistOf(in))
    })
  })
})
