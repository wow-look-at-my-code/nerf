package commands

import (
	"os"

	"path_prefix/src/common"
)

func init() {
	common.Register("bash", Bash)
}

func Bash() {
	// Unset CLAUDECODE so that commands within bash scripts
	// (like sed, find, etc.) use their real implementations
	// instead of being blocked by our wrappers.
	os.Unsetenv("CLAUDECODE")

	// Pass through to real bash
	common.ExecReal("bash", os.Args[1:])
}
