package options

import (
  "fmt"
  "io"
  "os"
  "path"

  "github.com/spf13/cobra"
  "github.com/spf13/pflag"
)

var _ Option = &Output{}
var _ PostRunOption = &Output{}

// Output is an option that adds mechanisms for writing output text (e.g. to a file)
type Output struct {
  // Out is an io.Writer, that can be used to write output to the given file.
  Out io.Writer

  outFile string
}

func (i *Output) AddFlags(fs *pflag.FlagSet) {
  fs.StringVarP(&i.outFile, "out", "o", "", "output file (- or default for stdin)")
}

func (i *Output) Complete(cmd *cobra.Command, args []string) error {
  if i.outFile == "" || i.outFile == "-" {
    i.Out = os.Stdout
    return nil
  }

  outputFile := i.outFile
  if !path.IsAbs(outputFile) {
    pwd, err := os.Getwd()
    if err != nil {
      return fmt.Errorf("failed to get current working directory: %w", err)
    }
    outputFile = path.Join(pwd, outputFile)
  }

  var err error
  i.Out, err = os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0600)
  if err != nil {
    return fmt.Errorf("failed to open output file: %w", err)
  }
  return nil
}

func (i *Output) PostRun(_ *cobra.Command, _ []string) error {
  if closer, ok := i.Out.(io.Closer); ok {
    if err := closer.Close(); err != nil {
      return fmt.Errorf("failed closing output file: %w", err)
    }
  }
  return nil
}
