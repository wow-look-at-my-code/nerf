package commands

import (
	"fmt"
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("fdfind", Fdfind)
}

func Fdfind() {
	cmd := "fd"
	if strings.HasSuffix(os.Args[0], "fdfind") {
		cmd = "fdfind"
	}

	// Block -x/-X (exec) flags
	for _, arg := range os.Args[1:] {
		if arg == "-x" || arg == "--exec" || arg == "-X" || arg == "--exec-batch" {
			fmt.Fprintf(os.Stderr, "Error: %s is not allowed in Claude Code.\n", arg)
			fmt.Fprintln(os.Stderr, "Use the Bash tool with explicit commands instead.")
			os.Exit(1)
		}
	}

	common.ExecReal(cmd, os.Args[1:])
}
