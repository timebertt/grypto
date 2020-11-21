package rsa32_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/onsi/gomega/gstruct"

  "github.com/timebertt/grypto/rsa32"
)

var _ = Describe("RSA32", func() {
  var (
    pub  *rsa32.PublicKey
    priv *rsa32.PrivateKey

    m, c int32
  )

  BeforeEach(func() {
    pub = &rsa32.PublicKey{N: 1271, E: 7}
    priv = &rsa32.PrivateKey{PublicKey: *pub, D: 343, P: 31, Q: 41}

    m = 42
    c = 1067
  })

  Describe("PublicKey", func() {
    Describe("#Encrypt", func() {
      It("should encrypt the message correctly", func() {
        Expect(pub.Encrypt(m)).To(Equal(c))
      })
    })
  })

  Describe("PrivateKey", func() {
    Describe("#Public", func() {
      It("should return the public key", func() {
        Expect(priv.Public()).To(PointTo(Equal(*pub)))
      })
    })

    Describe("#Decrypt", func() {
      It("should decrypt the message correctly", func() {
        Expect(priv.Decrypt(c)).To(Equal(m))
      })
    })
  })
})
