package options

import "github.com/spf13/cobra"

// AddEncryptDecryptSubcommands adds `encrypt` and `decrypt` subcommands to cmd.
func AddEncryptDecryptSubcommands(cmd *cobra.Command, cipherName string, runE func(cmd *cobra.Command, direction Direction, args []string) error) {
  cmd.AddCommand(
    &cobra.Command{
      Use:   "encrypt",
      Short: "encrypt plaintext using " + cipherName,
      RunE: func(cmd *cobra.Command, args []string) error {
        return runE(cmd, Encrypt, args)
      },
    },
    &cobra.Command{
      Use:   "decrypt",
      Short: "decrypt ciphertext using " + cipherName,
      RunE: func(cmd *cobra.Command, args []string) error {
        return runE(cmd, Decrypt, args)
      },
    },
  )
}
