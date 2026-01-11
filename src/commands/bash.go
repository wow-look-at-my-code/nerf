package commands

import (
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

	// Find -c flag and its command
	for i, arg := range args {
		if arg == "-c" && i+1 < len(args) {
			cmd := args[i+1]

			// Skip if already has safety options
			if strings.HasPrefix(cmd, "set -e") || strings.Contains(cmd, "set -euo pipefail") {
				common.ExecReal("bash", args)
				return
			}

			// Inject set -euo pipefail at the start
			args[i+1] = "set -euo pipefail; " + cmd
			common.ExecReal("bash", args)
			return
		}
	}

	// No -c flag, pass through
	common.ExecReal("bash", args)
}
