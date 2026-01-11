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

	// Check for login shell flag - skip injection as login shells read config files
	// that may not be compatible with strict mode
	for _, arg := range args {
		if arg == "-l" || arg == "--login" {
			common.ExecReal("zsh", args)
			return
		}
	}

	// Find -c flag and its command
	for i, arg := range args {
		if arg == "-c" && i+1 < len(args) {
			cmd := args[i+1]

			// Skip if already has safety options
			if strings.HasPrefix(cmd, "set -e") || strings.Contains(cmd, "set -euo pipefail") {
				common.ExecReal("zsh", args)
				return
			}

			// Inject set -euo pipefail at the start
			args[i+1] = "set -euo pipefail; " + cmd
			common.ExecReal("zsh", args)
			return
		}
	}

	// No -c flag, pass through
	common.ExecReal("zsh", args)
}
