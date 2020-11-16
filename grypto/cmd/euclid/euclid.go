package euclid

import (
  "fmt"
  "strconv"

  "github.com/spf13/cobra"

  "github.com/timebertt/grypto/euclid"
  "github.com/timebertt/grypto/internal/unicode"
)

func NewCommand() *cobra.Command {
  var a, b int

  cmd := &cobra.Command{
    Use:     "euclid [a] [b]",
    Aliases: []string{"gcd"},
    Short:   "Use Euclid's algorithm to calculate the greatest common divisor (gcd)",
    Long: `Euclid's algorithm is a fairly efficient algorithm for calculating the greatest common divisor (gcd)
of two integers. It is a  based on the following two cases:
  gcd(a, 0) = 0
  gcd(a, b) = gcd(b, a mod b)

In it not only calculates the gcd of two integers but additionally calculates two integers a and b, such that
  gcd(a, b) = x*a + y*b
If gcd(a, b) = 1, y is b's multiplicative inverse in ` + unicode.ZSubscriptSmallA + ` (y * b ` + unicode.IdenticalTo + ` 1 mod a).

See: https://en.wikipedia.org/wiki/Euclidean_algorithm, https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm`,
    Args: cobra.ExactArgs(2),
    PreRunE: func(cmd *cobra.Command, args []string) error {
      var err error
      a, err = strconv.Atoi(args[0])
      if err != nil {
        return fmt.Errorf("first argument is not an int: %w", err)
      }

      b, err = strconv.Atoi(args[1])
      if err != nil {
        return fmt.Errorf("second argument is not an int: %w", err)
      }

      cmd.SilenceUsage = true

      return nil
    },
    RunE: func(cmd *cobra.Command, args []string) error {
      return runEuclid(a, b)
    },
  }

  return cmd
}

func runEuclid(a, b int) (err error) {
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

  gcd, x, y := euclid.GreatestCommonDivisorExtended(a, b)

  fmt.Printf("gcd(%d,%d) = %d = %s*%d + %s*%d\n", a, b, gcd, parenthesis(x), a, parenthesis(y), b)

  if gcd == 1 {
    fmt.Printf("=> %d%s %s %d mod %d\n", b, unicode.SuperscriptMinusOne, unicode.IdenticalTo, y, a)
  }

  return nil
}

func parenthesis(i int) string {
  if i < 0 {
    return fmt.Sprintf("(%d)", i)
  }
  return fmt.Sprintf("%d", i)
}
