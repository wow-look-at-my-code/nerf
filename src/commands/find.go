package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"path_prefix/src/common"
)

func init() {
	common.Register("find", Find)
}

var dangerousFlags = []string{
	"-delete",
	"-exec",
	"-execdir",
	"-ok",
	"-okdir",
}

func Find() {
	// Check for dangerous flags
	for _, arg := range os.Args[1:] {
		argLower := strings.ToLower(arg)
		for _, dangerous := range dangerousFlags {
			if argLower == dangerous {
				fmt.Fprintf(os.Stderr, "Error: %s is not allowed in Claude Code.\n", arg)
				fmt.Fprintln(os.Stderr, "Use the Bash tool with explicit commands instead.")
				os.Exit(1)
			}
		}
	}

	// Remap find to fd with -uu (unrestricted)
	// fd uses different syntax, so we do a basic translation
	fdArgs := []string{"-uu"}

	// Simple translation: find [path] -name "pattern" -> fd "pattern" [path]
	var paths []string
	var pattern string
	skipNext := false

	for i, arg := range os.Args[1:] {
		if skipNext {
			skipNext = false
			continue
		}

		switch arg {
		case "-name", "-iname":
			if i+1 < len(os.Args[1:]) {
				pattern = os.Args[i+2]
				if arg == "-iname" {
					fdArgs = append(fdArgs, "-i")
				}
				skipNext = true
			}
		case "-type":
			if i+1 < len(os.Args[1:]) {
				t := os.Args[i+2]
				switch t {
				case "f":
					fdArgs = append(fdArgs, "-t", "f")
				case "d":
					fdArgs = append(fdArgs, "-t", "d")
				case "l":
					fdArgs = append(fdArgs, "-t", "l")
				}
				skipNext = true
			}
		case "-maxdepth":
			if i+1 < len(os.Args[1:]) {
				fdArgs = append(fdArgs, "-d", os.Args[i+2])
				skipNext = true
			}
		default:
			if !strings.HasPrefix(arg, "-") {
				paths = append(paths, arg)
			}
		}
	}

	// Build final args: fd [options] [pattern] [path...]
	if pattern != "" {
		fdArgs = append(fdArgs, pattern)
	}
	fdArgs = append(fdArgs, paths...)

	// Execute fd
	fdPath, err := exec.LookPath("fd")
	if err != nil {
		fdPath, err = exec.LookPath("fdfind")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: fd/fdfind not found")
			os.Exit(127)
		}
	}

	argv := append([]string{"fd"}, fdArgs...)
	syscall.Exec(fdPath, argv, os.Environ())
}
