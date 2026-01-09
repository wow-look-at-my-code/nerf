package main

import (
	"fmt"
	"os"

	"path_prefix/src/commands"
	"path_prefix/src/common"
)

func main() {
	cmd := common.Init()

	if !common.ShouldWrap() {
		common.ExecReal(cmd, os.Args[1:])
	}

	switch cmd {
	case "cat":
		commands.Cat()
	case "docker":
		commands.Docker()
	case "find":
		commands.Find()
	case "grep":
		commands.Grep()
	case "head":
		commands.Head()
	case "sed":
		commands.Sed()
	case "sleep":
		commands.Sleep()
	case "tail":
		commands.Tail()
	case "fd", "fdfind":
		commands.Fdfind()
	default:
		fmt.Fprintf(os.Stderr, "nerf: unknown command %q\n", cmd)
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Usage: Create a symlink with a supported command name pointing to this binary.")
		fmt.Fprintln(os.Stderr, "Example: ln -s /path/to/nerf /path/to/cat")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Supported commands: cat, docker, fd, fdfind, find, grep, head, sed, sleep, tail")
		os.Exit(1)
	}
}
