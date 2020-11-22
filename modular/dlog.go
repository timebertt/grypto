package modular

// DLogMaxMod is the maximum mod value accepted in DLog32.
const DLogMaxMod = 1 << 22

// DLog32 calculates the discrete logarithm of x to the given base and modulus by enumeration for int32 numbers.
// The discrete logarithm of a number x to the base of b modulo m is defined as the smallest number y,
// so that b^y ≡ x mod m. DLog32 is the inverse operation to Pow32.
// Enumeration is a very simple approach to calculate dlog. It calculates b^i for i=0,1,...,m until b^i=x.
// While being simple, the algorithm can take up to order(b) steps in the worst case, so it is very impractical
// for bases with large order.
// Calculating the discrete logarithm is thought to be hard, so currently there is no known algorithm for solving
// it efficiently. The security of some cryptographic algorithms (e.g. Diffie-Hellman, ElGamal and others) is based
// on exactly this assumption, that DLog32 is hard.
// See https://en.wikipedia.org/wiki/Discrete_logarithm.
func DLog32(x, base, mod int32) (dlog int32, exists bool) {
  if x <= 0 {
    panic("grypt/modular: x must be greater than 0")
  }
  if base <= 0 {
    panic("grypt/modular: base must be greater than 0")
  }
  if mod <= 0 {
    panic("grypto/modular: modulus must be greater than 0")
  }
  if mod > DLogMaxMod {
    panic("grypto/modular: modulus too large")
  }

  var (
    x64 = int64(x)
    exp = int64(1)
    b   = int64(base % mod)
    m   = int64(mod)
  )

  dlog = 0
  for ; dlog < mod; dlog++ {
    if exp == x64 {
      return dlog, true
    }

    if dlog > 0 && exp == 1 {
      break
    }
    if exp == 0 {
      return 0, false
    }

    exp = exp * b % m
  }

  return 0, false
}
