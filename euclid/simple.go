package euclid

// GreatestCommonDivisor calculates the greatest common divisor (gcd) of two integers using the simple form of Euclid's
// algorithm. It is a fairly efficient algorithm based on the following two cases:
//   gcd(a, 0) = 0
//   gcd(a, b) = gcd(b, a mod b)
// See: https://en.wikipedia.org/wiki/Euclidean_algorithm
func GreatestCommonDivisor(a, b int) int {
  return euclid(a, b)
}

func euclid(a, b int) int {
  for b != 0 {
    a, b = b, a%b
  }
  return abs(a)
}

func abs(i int) int {
  if i < 0 {
    return -i
  }
  return i
}
