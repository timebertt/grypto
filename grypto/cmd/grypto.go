/*
Copyright © 2020 Tim Ebert

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
  "github.com/spf13/cobra"

  "github.com/timebertt/grypto/grypto/cmd/caesar"
  "github.com/timebertt/grypto/grypto/cmd/euclid"
  "github.com/timebertt/grypto/grypto/cmd/exp"
)

func NewGryptoCommand() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "grypto",
    Short: "A collection of cryptographic algorithms implemented in go",
    Long: `grypto is a collection of cryptographic algorithms implemented in go.

It was implemented by Tim Ebert as a practical exercise to understand the fundamental mathematical
concepts behind cryptographic algorithms that were discussed in a lecture on cryptography in his
Computer Science Master studies.

The grypto CLI can be used to test and demonstrate the different algorithms implemented in the
grypto library.

WARNING: Please use this only for learning purposes! The grypto library and CLI were only build to
demonstrate and understand the basics of different cryptographic algorithms. There is no guarantee
on correctness, security and quality of the implementation. The implementation might not be compatible
with proper implementations of the different algorithms and might be vulnerable to attacks. Please use
the respective official implementations of the Go standard library (see https://golang.org/pkg/crypto/)
for writing secure Go applications.`,
  }

  cmd.AddCommand(
    caesar.NewCommand(),
    exp.NewCommand(),
    euclid.NewCommand(),
  )

  return cmd
}
