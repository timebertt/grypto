package options

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Option is a command line option that can add some flags and be completed.
type Option interface {
	AddFlags(fs *pflag.FlagSet)
	Complete(cmd *cobra.Command, args []string) error
}

// PostRunOption is a command line option that has an action associated with it, that should be run after the command
// has finished.
type PostRunOption interface {
	PostRun(cmd *cobra.Command, args []string) error
}

// Direction decides if the user wants to use a given cipher either for encryption or decryption.
type Direction bool

const (
	Encrypt = Direction(false)
	Decrypt = Direction(true)
)
