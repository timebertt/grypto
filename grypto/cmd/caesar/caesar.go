package caesar

import (
  "bytes"
  "crypto/cipher"
  "fmt"
  "io"
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
    input  = &options.Input{}
    key    = &options.Key{}
    output = &options.Output{}

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
      if err := output.Complete(cmd, args); err != nil {
        return err
      }

      keyInput := &bytes.Buffer{}
      if _, err := io.Copy(keyInput, key.In); err != nil {
        return fmt.Errorf("error reading key: %w", err)
      }
      if keyInput.Len() == 0 {
        return fmt.Errorf("given key is empty")
      }

      var err error
      parsedKey, err = strconv.Atoi(strings.TrimSuffix(keyInput.String(), "\n"))
      if err != nil {
        return fmt.Errorf("given key is not an int: %w", err)
      }

      cmd.SilenceErrors = true
      cmd.SilenceUsage = true
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
    return run(direction, parsedKey, input.In, output.Out)
  })

  options.AddAllFlags(cmd.PersistentFlags(),
    input,
    output,
    key,
  )

  return cmd
}

func run(direction options.Direction, key int, in io.Reader, out io.Writer) (err error) {
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

  _, err = io.Copy(out, block.NewBlockModeReader(blockMode, in))
  if err != nil {
    return fmt.Errorf("error crypting with %s: %w", CipherName, err)
  }

  return nil
}
