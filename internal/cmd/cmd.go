// command line handler
package cmd

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/raymond-design/go-emoji-func/internal/tracing"
	"github.com/raymond-design/go-emoji-func/internal/ver"
)

// Main command line implementation
func Run(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	gef := &cobra.Command{Use: "gef", SilenceUsage: true}

	gef.AddCommand(getVersion)

	gef.SetArgs(args)
	gef.SetIn(stdin)
	gef.SetOut(stdout)
	gef.SetErr(stderr)

	// running cli
	cont, clean, err := tracing.Trace(context.Background())
	
	if err != nil {
		fmt.Printf("couldn't trace: %v\n", err)
		return 1
	}
	defer clean()

	if err := gef.ExecuteContext(ctx); err == nil {
		return 0
	}

	if exit, ok := err.(*exec.ExitError); ok {
		return exit.ExitCode()
	}
	return 1
}

var version string

// get version info
var getVersion = &cobra.Command{
	Use: "version",
	Short: "Show gef version",
	Run: func(cmd *cobra.Command, args []string) {
		if version == "" {
			fmt.Printf("Go Emoji Func %s\n", ver.Version)
		} else {
			fmt.Printf("Go Emoji Fund %s\n", version)
		}
	},
}