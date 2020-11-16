package des

import (
  "crypto/cipher"
  "encoding/base64"
  "fmt"
  "io"

  "github.com/spf13/cobra"

  "github.com/timebertt/grypto/block"
  "github.com/timebertt/grypto/des"
  "github.com/timebertt/grypto/grypto/options"
)

const (
  CipherName = "DES"
)

func NewCommand() *cobra.Command {
  var (
    input   = &options.Input{}
    keyFlag = &options.Key{}
    output  = &options.Output{}

    key = make([]byte, 8)
  )

  cmd := &cobra.Command{
    Use:   "des",
    Short: "Use " + CipherName + " for encryption and decryption",
    Long:  CipherName + ` (Data Encryption Standard)`,
    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
      if err := input.Complete(cmd, args); err != nil {
        return err
      }
      if err := keyFlag.Complete(cmd, args); err != nil {
        return err
      }
      if err := output.Complete(cmd, args); err != nil {
        return err
      }

      if _, err := io.ReadFull(base64.NewDecoder(base64.StdEncoding, keyFlag.In), key); err != nil {
        return fmt.Errorf("error reading key: %w", err)
      }

      cmd.SilenceErrors = true
      cmd.SilenceUsage = true
      return nil
    },
    PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
      if err := input.PostRun(cmd, args); err != nil {
        return err
      }
      if err := keyFlag.PostRun(cmd, args); err != nil {
        return err
      }
      return nil
    },
  }

  options.AddEncryptDecryptSubcommands(cmd, CipherName, func(cmd *cobra.Command, direction options.Direction, args []string) error {
    return run(direction, key, input.In, output.Out)
  })

  options.AddAllFlags(cmd.PersistentFlags(),
    input,
    output,
    keyFlag,
  )

  return cmd
}

func run(direction options.Direction, key []byte, in io.Reader, out io.Writer) (err error) {
  defer func() {
    if p := recover(); p != nil {
      if e, ok := p.(error); ok {
        err = e
      }
    }
  }()

  d, err := des.NewCipher(key)
  if err != nil {
    return err
  }

  var blockMode cipher.BlockMode
  if direction == options.Decrypt {
    blockMode = block.NewECBDecrypter(d)
  } else {
    blockMode = block.NewECBEncrypter(d)
  }

  wr := base64.NewEncoder(base64.StdEncoding, out)
  _, err = io.Copy(wr, block.NewBlockModeReader(blockMode, base64.NewDecoder(base64.StdEncoding, in)))
  if err != nil {
    return fmt.Errorf("error crypting with %s: %w", CipherName, err)
  }
  err = wr.Close()
  if err != nil {
    return fmt.Errorf("error crypting with %s: %w", CipherName, err)
  }

  return nil
}
