package options

import (
  "bytes"
  "fmt"
  "io"
  "os"
  "path"

  "github.com/spf13/cobra"
  "github.com/spf13/pflag"
)

var _ Option = &Input{}
var _ PostRunOption = &Input{}

// Input is an option that adds mechanisms for reading input text (either from a file or from a flag value)
type Input struct {
  // In is an io.Reader, that can be used to read the given input text from.
  In io.Reader

  inFile, inText string
}

func (i *Input) AddFlags(fs *pflag.FlagSet) {
  fs.StringVarP(&i.inFile, "in", "i", "", "input file for cipher (- for stdin)")
  fs.StringVarP(&i.inText, "in-text", "I", "", "input text for cipher")
}

func (i *Input) Complete(cmd *cobra.Command, args []string) error {
  if i.inFile == "" && i.inText == "" {
    return fmt.Errorf("no input was specified, input must be given either via file ('--in') or directly " +
      "as a flag value ('--in-text')")
  }

  if i.inFile != "" {
    inputFile := i.inFile
    if inputFile == "-" {
      i.In = os.Stdin
      return nil
    }

    if !path.IsAbs(inputFile) {
      pwd, err := os.Getwd()
      if err != nil {
        return fmt.Errorf("failed to get current working directory: %w", err)
      }
      inputFile = path.Join(pwd, inputFile)
    }

    var err error
    i.In, err = os.Open(inputFile)
    if err != nil {
      return fmt.Errorf("failed to open input file: %w", err)
    }
    return nil
  }

  i.In = bytes.NewBufferString(i.inText)
  return nil
}

func (i *Input) PostRun(_ *cobra.Command, _ []string) error {
  if closer, ok := i.In.(io.Closer); ok {
    if err := closer.Close(); err != nil {
      return fmt.Errorf("failed closing input file: %w", err)
    }
  }
  return nil
}
