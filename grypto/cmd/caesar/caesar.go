package caesar

import (
  "bytes"
  "crypto/cipher"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"

  "github.com/spf13/cobra"

  "github.com/timebertt/grypto/block"
  "github.com/timebertt/grypto/caesar"
  "github.com/timebertt/grypto/grypto/options"
)

const (
  CipherName = "Caesar Cipher"
)

func NewCommand() *cobra.Command {
  var (
    input = &options.Input{}
    key   = &options.Key{}

    parsedKey int
  )

  cmd := &cobra.Command{
    Use:   "caesar",
    Short: "Use the " + CipherName + " for encryption and decryption",
    Long: CipherName + ` is is a substituting block cipher operating on blocks of length 1 (single bytes).
Latin characters are encrypted by replacing them by the key-th next character in the alphabet.
The characters' case is kept and non-latin characters are not replaced.
See: https://en.wikipedia.org/wiki/Caesar_cipher`,
    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
      if err := input.Complete(cmd, args); err != nil {
        return err
      }
      if err := key.Complete(cmd, args); err != nil {
        return err
      }

      keyInput := &bytes.Buffer{}
      if _, err := io.Copy(keyInput, key.In); err != nil {
        return fmt.Errorf("error reading input: %w", err)
      }
      if keyInput.Len() == 0 {
        return fmt.Errorf("given key is empty")
      }

      var err error
      parsedKey, err = strconv.Atoi(strings.TrimSuffix(keyInput.String(), "\n"))
      if err != nil {
        return fmt.Errorf("given key is not an int: %w", err)
      }

      return nil
    },
    PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
      if err := input.PostRun(cmd, args); err != nil {
        return err
      }
      if err := key.PostRun(cmd, args); err != nil {
        return err
      }
      return nil
    },
  }

  options.AddEncryptDecryptSubcommands(cmd, CipherName, func(cmd *cobra.Command, direction options.Direction, args []string) error {
    return runCaesar(direction, parsedKey, input.In)
  })

  input.AddFlags(cmd.PersistentFlags())
  key.AddFlags(cmd.PersistentFlags())

  return cmd
}

func runCaesar(direction options.Direction, key int, input io.Reader) (err error) {
  defer func() {
    if p := recover(); p != nil {
      if e, ok := p.(error); ok {
        err = e
      }
    }
  }()

  var blockMode cipher.BlockMode
  if direction == options.Decrypt {
    blockMode = block.NewECBDecrypter(caesar.NewCipher(key))
  } else {
    blockMode = block.NewECBEncrypter(caesar.NewCipher(key))
  }

  return block.CopyBlockMode(blockMode, os.Stdout, input)
}
