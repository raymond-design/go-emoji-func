// command line interface for the emoji parser
package cli

import (
	"os"

	"github.com/raymond-design/go-emoji-func/internal/cmd"
)

// implementation to handle args is in internal/cmd
func Run(args []string) int {
	return cmd.Do(args, os.Stdin, os.Stdout, os.Stderr)
}