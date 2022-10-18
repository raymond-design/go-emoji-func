// interface for usage within code
package main

import (
	"os"

	"github.com/raymond-design/go-emoji-func/internal/cmd"
)

func AllowEmojiFunc() {
	os.Exit(cmd.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}