package commands

import (
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("zsh", Zsh)
}

// Zsh wraps zsh to inject set -euo pipefail for -c commands
func Zsh() {
	args := os.Args[1:]

	// Find -c flag and its command
	for i, arg := range args {
		if arg == "-c" && i+1 < len(args) {
			cmd := args[i+1]

			// Skip if already has safety options
			if strings.HasPrefix(cmd, "set -e") || strings.Contains(cmd, "set -euo pipefail") {
				common.ExecReal("zsh", args)
				return // ExecReal replaces process, but return for safety
			}

			// Inject set -euo pipefail at the start
			args[i+1] = "set -euo pipefail; " + cmd
			common.ExecReal("zsh", args)
			return // ExecReal replaces process, but return for safety
		}
	}

	// No -c flag, pass through
	common.ExecReal("zsh", args)
}
