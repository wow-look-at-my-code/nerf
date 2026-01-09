package commands

import (
	"os"
	"strings"

	"path_prefix/src/common"
)

func Docker() {
	args := os.Args[1:]

	// Add quiet flags to commands that support them
	if len(args) >= 2 {
		joined := strings.Join(args[:2], " ")
		var flag string
		switch joined {
		case "compose up":
			flag = "--quiet-pull"
		case "compose pull":
			flag = "--quiet"
		}
		if flag != "" {
			newArgs := make([]string, 0, len(args)+1)
			newArgs = append(newArgs, args[:2]...)
			newArgs = append(newArgs, flag)
			newArgs = append(newArgs, args[2:]...)
			args = newArgs
		}
	}
	if len(args) >= 1 && args[0] == "build" {
		newArgs := make([]string, 0, len(args)+1)
		newArgs = append(newArgs, args[0])
		newArgs = append(newArgs, "--progress=quiet")
		newArgs = append(newArgs, args[1:]...)
		args = newArgs
	}
	common.ExecReal("docker", args)
}
