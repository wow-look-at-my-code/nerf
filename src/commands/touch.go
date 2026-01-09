package commands

import (
	"fmt"
	"os"

	"path_prefix/src/common"
)

func init() {
	common.Register("touch", Touch)
}

func Touch() {
	args := os.Args[1:]

	// Check each file argument - block if file already exists
	for _, arg := range args {
		if arg[0] == '-' {
			continue // skip flags
		}
		if _, err := os.Stat(arg); err == nil {
			fmt.Fprintf(os.Stderr, "Error: touch is not allowed on existing files in Claude Code.\n")
			fmt.Fprintf(os.Stderr, "File already exists: %s\n", arg)
			fmt.Fprintln(os.Stderr, "")
			fmt.Fprintln(os.Stderr, "touch can only be used to create new files:")
			fmt.Fprintln(os.Stderr, "  touch newfile.txt        # OK - creates new file")
			fmt.Fprintln(os.Stderr, "  touch existing.txt       # NOT ALLOWED")
			os.Exit(1)
		}
	}

	// All files are new, pass through to real touch
	common.ExecReal("touch", args)
}
