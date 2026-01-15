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

	if !common.ShouldWrap() {
		common.ExecReal(cmd, os.Args[1:])
		return
	}

	if handler, ok := common.Handlers[cmd]; ok {
		if handler().IsHandled() {
			return
		}
		// Handler returned PassThru, pass through to real command
		common.ExecReal(cmd, os.Args[1:])
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
