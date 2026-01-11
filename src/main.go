package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	_ "path_prefix/src/commands"
	"path_prefix/src/common"
)

func main() {
	cmd := common.Init()

	fmt.Fprintf(os.Stderr, "DEBUG main: cmd=%q CLAUDECODE=%q\n", cmd, os.Getenv("CLAUDECODE"))

	if !common.ShouldWrap() {
		fmt.Fprintf(os.Stderr, "DEBUG main: passthrough (no CLAUDECODE)\n")
		common.ExecReal(cmd, os.Args[1:])
		return // ExecReal replaces process, but return for safety
	}

	if handler, ok := common.Handlers[cmd]; ok {
		fmt.Fprintf(os.Stderr, "DEBUG main: calling handler for %q\n", cmd)
		handler()
		return
	}

	// Build sorted list of supported commands
	cmds := make([]string, 0, len(common.Handlers))
	for k := range common.Handlers {
		cmds = append(cmds, k)
	}
	sort.Strings(cmds)

	fmt.Fprintf(os.Stderr, "nerf: unknown command %q\n", cmd)
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage: Create a symlink with a supported command name pointing to this binary.")
	fmt.Fprintln(os.Stderr, "Example: ln -s /path/to/nerf /path/to/cat")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintf(os.Stderr, "Supported commands: %s\n", strings.Join(cmds, ", "))
	os.Exit(1)
}
