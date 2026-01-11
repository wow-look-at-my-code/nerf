package commands

import (
	"fmt"
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("docker", Docker)
}

func restartError() {
	fmt.Fprintln(os.Stderr, "Error: 'docker compose restart' doesn't reload config changes or rebuild images.")
	fmt.Fprintln(os.Stderr, "Use 'docker compose up -d' instead to recreate containers with new config.")
	os.Exit(1)
}

// findSubcommand finds the first non-flag argument
func findSubcommand(args []string) string {
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

func Docker() {
	args := os.Args[1:]

	// DEBUG
	fmt.Fprintf(os.Stderr, "DEBUG docker: args=%v\n", args)

	// Check for compose subcommand
	if len(args) >= 1 && args[0] == "compose" {
		subcmd := findSubcommand(args[1:])
		fmt.Fprintf(os.Stderr, "DEBUG docker: subcmd=%q\n", subcmd)

		if subcmd == "restart" {
			restartError()
		}

		// Add quiet flags (only for simple cases without intermediate flags)
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
	}

	if len(args) >= 1 && args[0] == "build" {
		newArgs := make([]string, 0, len(args)+1)
		newArgs = append(newArgs, args[0])
		newArgs = append(newArgs, "--progress=quiet")
		// Filter out --no-cache flag
		for _, arg := range args[1:] {
			if arg != "--no-cache" {
				newArgs = append(newArgs, arg)
			}
		}
		args = newArgs
	}

	common.ExecReal("docker", args)
}
