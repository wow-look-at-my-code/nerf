package commands

import (
	"fmt"
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("sed", Sed)
}

func Sed() common.HandlerResult {
	// Check for -i flag (in-place editing) - always block
	for _, arg := range os.Args[1:] {
		if arg == "-i" || strings.HasPrefix(arg, "-i") {
			blockSed()
		}
	}

	// If file arguments provided, block (should use Read/Edit tools instead)
	if common.HasFileArgs(os.Args[1:]) {
		blockSed()
	}

	// Piped input is allowed - pass through to real sed
	return common.PassThru
}

func blockSed() {
	fmt.Fprintln(os.Stderr, "Error: sed is not allowed for in-place file editing in Claude Code.")
	fmt.Fprintln(os.Stderr, "Use the Edit tool instead to modify files.")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "sed can only be used for stream processing (piped input):")
	fmt.Fprintln(os.Stderr, "  echo 'text' | sed 's/old/new/'   # OK")
	fmt.Fprintln(os.Stderr, "  sed -i 's/old/new/' file.txt     # NOT ALLOWED")
	os.Exit(1)
}
