package commands

import (
	"fmt"
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("gh", Gh)
}

func Gh() {
	args := os.Args[1:]

	// Block "gh run watch" - should use gh wait-ci instead
	if len(args) >= 2 {
		joined := strings.Join(args[:2], " ")
		if joined == "run watch" {
			fmt.Fprintln(os.Stderr, "Error: Use 'gh wait-ci' instead of 'gh run watch'")
			os.Exit(1)
		}
	}

	common.ExecReal("gh", args)
}
