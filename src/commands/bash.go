package commands

import (
	"os"

	"path_prefix/src/common"
)

func init() {
	common.Register("bash", Bash)
}

func Bash() {
	// Set NERF_IN_SCRIPT so that commands within bash scripts
	// (like sed, find, etc.) use their real implementations.
	// ShouldWrap() checks for this variable and passes through when set.
	os.Setenv("NERF_IN_SCRIPT", "1")

	// Pass through to real bash
	common.ExecReal("bash", os.Args[1:])
}
