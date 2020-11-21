package rsa32

import (
  "github.com/timebertt/grypto/modular"
)

func GenerateKeyPair() (*PublicKey, *PrivateKey, error) {
  panic("implement me")
}

type PublicKey struct {
  N int32 // modulus (N = P*Q)
  E int32 // public exponent for encryption
}

func (pub *PublicKey) Encrypt(msg int32) (int32, error) {
  return modular.Pow32(msg, pub.E, pub.N), nil
}

type PrivateKey struct {
  PublicKey // public part needed for decryption (e.g. modulus)

  D    int32 // private exponent for decryption (1 ≡ D*E mod (P-1)*(Q-1))
  P, Q int32 // primes (must be kept secret, otherwise attackers can calculate D quite easily)
}

func (priv *PrivateKey) Public() *PublicKey {
  return &priv.PublicKey
}

func (priv *PrivateKey) Decrypt(msg int32) (int32, error) {
  return modular.Pow32(msg, priv.D, priv.N), nil
}
