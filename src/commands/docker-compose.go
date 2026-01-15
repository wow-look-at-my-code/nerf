package commands

import (
	"fmt"
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("docker-compose", DockerCompose)
}

func dockerComposeRestartError() {
	fmt.Fprintln(os.Stderr, "Error: 'docker-compose restart' doesn't reload config changes or rebuild images.")
	fmt.Fprintln(os.Stderr, "Use 'docker-compose up -d' instead to recreate containers with new config.")
	os.Exit(1)
}

// findDockerComposeSubcommand finds the first non-flag argument
func findDockerComposeSubcommand(args []string) string {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			// Skip flags that take values
			if arg == "--ansi" || arg == "-f" || arg == "--file" || arg == "-p" || arg == "--project-name" {
				i++
			}
			continue
		}
		return arg
	}
	return ""
}

// DockerCompose handles the standalone docker-compose command
func DockerCompose() common.HandlerResult {
	args := os.Args[1:]

	subcmd := findDockerComposeSubcommand(args)

	if subcmd == "restart" {
		dockerComposeRestartError()
	}

	// Add quiet flags (only for simple cases without intermediate flags)
	if len(args) >= 1 {
		var flag string
		switch args[0] {
		case "up":
			flag = "--quiet-pull"
		case "pull":
			flag = "--quiet"
		case "build":
			flag = "--progress=quiet"
		}
		if flag != "" {
			newArgs := make([]string, 0, len(args)+1)
			newArgs = append(newArgs, args[0])
			newArgs = append(newArgs, flag)
			newArgs = append(newArgs, args[1:]...)
			args = newArgs
		}
	}

	common.ExecReal("docker-compose", args)
	return common.Handled
}
