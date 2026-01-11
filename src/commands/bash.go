package commands

import (
	"fmt"
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("bash", Bash)
}

// Bash wraps bash to inject set -euo pipefail for -c commands
func Bash() {
	args := os.Args[1:]

	fmt.Fprintf(os.Stderr, "DEBUG bash: args=%v\n", args)

	// Find -c flag and its command
	for i, arg := range args {
		if arg == "-c" && i+1 < len(args) {
			cmd := args[i+1]

			// Skip if already has safety options
			if strings.HasPrefix(cmd, "set -e") || strings.Contains(cmd, "set -euo pipefail") {
				fmt.Fprintf(os.Stderr, "DEBUG bash: skipping injection (already has set -e)\n")
				common.ExecReal("bash", args)
				return // ExecReal replaces process, but return for safety
			}

			// Inject set -euo pipefail at the start
			args[i+1] = "set -euo pipefail; " + cmd
			fmt.Fprintf(os.Stderr, "DEBUG bash: injected, new cmd=%q\n", args[i+1])
			common.ExecReal("bash", args)
			return // ExecReal replaces process, but return for safety
		}
	}

	// No -c flag, pass through
	fmt.Fprintf(os.Stderr, "DEBUG bash: no -c flag, passthrough\n")
	common.ExecReal("bash", args)
}
