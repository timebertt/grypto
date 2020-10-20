package euclid

// GreatestCommonDivisorExtended also calculates the greatest common divisor (gcd) of two integers but additionally
// calculates two integers a and b, such that
//   gcd(a, b) = x*a + y*b
// If gcd(a, b) = 1, y is b's multiplicative inverse in ℤₐ (y * b ≡ 1 mod a).
// See: https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
func GreatestCommonDivisorExtended(a, b int) (gcd, x, y int) {
  if a < 0 || b < 0 {
    panic("input may not be negative")
  }
  return extendedEuclid(a, b)
}

func extendedEuclid(a, b int) (gcd, x, y int) {
  var (
    x0, x1 = 1, 0
    y0, y1 = 0, 1
    q      int
  )

  for b != 0 {
    q = a / b
    a, b = b, a%b
    x0, x1 = x1, x0-q*x1
    y0, y1 = y1, y0-q*y1
  }
  return abs(a), x0, y0
}
