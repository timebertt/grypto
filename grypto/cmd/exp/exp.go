package exp

import (
  "fmt"
  "math"
  "strconv"

  "github.com/spf13/cobra"

  "github.com/timebertt/grypto/modular"
)

func NewCommand() *cobra.Command {
  var base, exp, mod int32

  cmd := &cobra.Command{
    Use:     "exp [base] [exponent] [modulus]",
    Aliases: []string{"mod-exp", "square-and-multiply"},
    Short:   "Use the square-and-multiply method for calculating the modular exponentiation",
    Long: `The exp command implements modular exponentiation using the square-and-multiply method for int32 numbers.
It prints a value x so that x = base ^ exp mod m.

Modular exponentiation is heavily used e.g. for primality tests and public-key cryptography (like RSA).
Even for reasonably small integers, calculating the modular exponentiation directly is on the one hand
quite inefficient and on the other hand very impractical, as the resulting integers will easily outgrow
the usual variable/register sizes.

A fairly efficient method is exponentiation by squaring (also known as square-and-multiply or binary
exponentiation). It calculates the modular squares of base and multiplies all squares for which the exp
has a 1 in its binary notation.
See https://en.wikipedia.org/wiki/Exponentiation_by_squaring.`,
    Args: cobra.ExactArgs(3),
    PreRunE: func(cmd *cobra.Command, args []string) error {
      b, err := strconv.Atoi(args[0])
      if err != nil {
        return fmt.Errorf("first argument is not an int: %w", err)
      }
      if b > math.MaxInt32 {
        return fmt.Errorf("base is greater than MaxInt32 (%d): %d", math.MaxInt32, b)
      }
      base = int32(b)

      e, err := strconv.Atoi(args[1])
      if err != nil {
        return fmt.Errorf("second argument is not an int: %w", err)
      }
      if e > math.MaxInt32 {
        return fmt.Errorf("exponent is greater than MaxInt32 (%d): %d", math.MaxInt32, b)
      }
      exp = int32(e)

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
      return runPow32(base, exp, mod)
    },
  }

  return cmd
}

func runPow32(base, exp, mod int32) (err error) {
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

  pow := modular.Pow32(base, exp, mod)

  fmt.Printf("%s ^ %s mod %s = %d\n", parenthesis(base), parenthesis(exp), parenthesis(mod), pow)
  return nil
}

func parenthesis(i int32) string {
  if i < 0 {
    return fmt.Sprintf("(%d)", i)
  }
  return fmt.Sprintf("%d", i)
}
