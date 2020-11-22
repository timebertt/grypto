package dlog

import (
  "fmt"
  "math"
  "strconv"

  "github.com/spf13/cobra"

  "github.com/timebertt/grypto/modular"
)

func NewCommand() *cobra.Command {
  var x, base, mod int32

  cmd := &cobra.Command{
    Use:   "dlog [x] [base] [modulus]",
    Short: "Calculate the discrete logarithm of x to the given base mod modulus",
    Long: `dlog calculates the discrete logarithm of x to the given base and modulus by enumeration for int32 numbers.
The discrete logarithm of a number x to the base of b modulo m is defined as the smallest number y,
so that b^y ≡ x mod m. dlog is the inverse operation to exp.

Enumeration is a very simple approach to calculate dlog. It calculates b^i for i=0,1,...,m until b^i=x.
While being simple, the algorithm can take up to order(b) steps in the worst case, so it is very impractical
for bases with large order.

Calculating the discrete logarithm is thought to be hard, so currently there is no known algorithm for solving
it efficiently. The security of some cryptographic algorithms (e.g. Diffie-Hellman, ElGamal and others) is based
on exactly this assumption, that DLog is hard.
See https://en.wikipedia.org/wiki/Discrete_logarithm.`,
    Args: cobra.ExactArgs(3),
    PreRunE: func(cmd *cobra.Command, args []string) error {
      xIn, err := strconv.Atoi(args[0])
      if err != nil {
        return fmt.Errorf("first argument is not an int: %w", err)
      }
      if xIn > math.MaxInt32 {
        return fmt.Errorf("x is greater than MaxInt32 (%d): %d", math.MaxInt32, xIn)
      }
      x = int32(xIn)

      b, err := strconv.Atoi(args[1])
      if err != nil {
        return fmt.Errorf("second argument is not an int: %w", err)
      }
      if b > math.MaxInt32 {
        return fmt.Errorf("base is greater than MaxInt32 (%d): %d", math.MaxInt32, b)
      }
      base = int32(b)

      m, err := strconv.Atoi(args[2])
      if err != nil {
        return fmt.Errorf("third argument is not an int: %w", err)
      }
      if m > math.MaxInt32 {
        return fmt.Errorf("modulus is greater than MaxInt32 (%d): %d", math.MaxInt32, b)
      }
      mod = int32(m)

      cmd.SilenceErrors = true
      cmd.SilenceUsage = true

      return nil
    },
    RunE: func(cmd *cobra.Command, args []string) error {
      return runDLog32(x, base, mod)
    },
  }

  return cmd
}

func runDLog32(x, base, mod int32) (err error) {
  defer func() {
    if p := recover(); p != nil {
      if e, ok := p.(error); ok {
        err = e
      }
      if e, ok := p.(string); ok {
        err = fmt.Errorf(e)
      }
    }
  }()

  dlog, exists := modular.DLog32(x, base, mod)
  if !exists {
    return fmt.Errorf("dlog(%d) to the base %d mod %d does not exist\n", x, base, mod)
  }

  fmt.Printf("dlog(%d) to the base %d mod %d = %d\n", x, base, mod, dlog)

  return nil
}
