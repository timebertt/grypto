package modular

// Pow32 implements modular exponentiation using the square-and-multiply method for int32 numbers.
// It returns a value x so that x = base ^ exp mod m.
// Modular exponentiation is heavily used e.g. for primality tests and public-key cryptography (like RSA).
// Even for reasonably small integers, calculating the modular exponentiation directly is on the one hand
// quite inefficient and on the other hand very impractical, as the resulting integers will easily outgrow
// the usual variable/register sizes.
// A fairly efficient method is exponentiation by squaring (also known as square-and-multiply or binary
// exponentiation). It calculates the modular squares of base and multiplies all squares for which the exp
// has a 1 in its binary notation (see https://en.wikipedia.org/wiki/Exponentiation_by_squaring).
func Pow32(base, exp, mod int32) int32 {
  if mod <= 0 {
    panic("grypto/modular: modulus must be greater than 0")
  }
  if exp < 0 {
    panic("grypto/modular: negative exponent not allowed")
  }
  if base == 0 && exp == 0 {
    panic("grypto/modular: 0^0 is not defined")
  }

  // special cases
  if base == 0 {
    return 0
  }
  if exp == 0 {
    return 1
  }

  // normalize base
  if base < 0 {
    base += mod
  }

  // square-and-multiply (fast exponentiation)
  var (
    // calculations based on int64 to avoid overflows
    // (MaxInt32 < MaxInt32*MaxInt32 < MaxInt64)
    m = int64(mod)
    x = int64(1)
    a = int64(base)
  )

  for exp > 0 {
    if exp&1 == 1 {
      x = x * a % m // multiply
    }

    a = a * a % m // square
    exp >>= 1
  }

  // result will never outgrow modulus => x < mod <= MaxInt32
  return int32(x)
}
