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

var _ Option = &Key{}
var _ PostRunOption = &Key{}

// Key is an option that adds mechanisms for reading key text (either from a file or from a flag value)
type Key struct {
	// In is an io.Reader, that can be used to read the given key text from.
	In io.Reader

	keyFile, keyText string
}

func (i *Key) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&i.keyFile, "key", "k", "", "key file for cipher")
	fs.StringVarP(&i.keyText, "key-text", "K", "", "key text for cipher")
}

func (i *Key) Complete(cmd *cobra.Command, args []string) error {
	if i.keyFile == "" && i.keyText == "" {
		return fmt.Errorf("no key was specified, key must be given either via file ('--key') or directly " +
			"as a flag value ('--key-text')")
	}

	if i.keyFile != "" {
		var err error

		keyFile := i.keyFile
		if !path.IsAbs(keyFile) {
			pwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current working directory: %w", err)
			}
			keyFile = path.Join(pwd, keyFile)
		}

		i.In, err = os.Open(keyFile)
		if err != nil {
			return fmt.Errorf("failed to open key file: %w", err)
		}
		return nil
	}

	i.In = bytes.NewBufferString(i.keyText)
	return nil
}

func (i *Key) PostRun(_ *cobra.Command, _ []string) error {
	if closer, ok := i.In.(io.Closer); ok {
		if err := closer.Close(); err != nil {
			return fmt.Errorf("failed closing key file: %w", err)
		}
	}
	return nil
}
