package main

import (
	"os"
	"os/exec"
	"strings"
)

// Fake claude executable for testing - spawns bash which runs the command.
// This creates the process tree: command → bash → claude
// which is what ShouldWrap() checks for.
func main() {
	// Only run in Docker to prevent accidental use outside tests
	if _, err := os.Stat("/.dockerenv"); err != nil {
		os.Stderr.WriteString("fake claude: only runs in Docker\n")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		os.Exit(0)
	}

	// Build command string and run via bash -c
	// This creates: command → bash → claude
	cmdStr := strings.Join(os.Args[1:], " ")
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}
