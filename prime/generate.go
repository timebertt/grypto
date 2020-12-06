package prime

import (
  "math/rand"
  "time"
)

const maxGenRounds = 2 << 30

func Generate31() int32 {
  return Generate31n(31)
}

func Generate31n(bits int) int32 {
  if bits > 31 {
    panic("grypto/prime: can only generate primes with up to 31 bits")
  }
  rand.Seed(time.Now().UnixNano())

  var n int32 = 1<<bits - 1
  for i := 0; i < maxGenRounds; i++ {
    p := rand.Int31n(n)
    if IsPrime32(p) {
      return p
    }
  }
  panic("grypto/prime: no prime found in maximum number of rounds")
}
