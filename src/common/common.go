package common

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// Init filters our directory out of PATH and returns the command name from argv[0]
func Init() string {
	// Get the directory containing this binary
	execPath, err := os.Executable()
	if err != nil {
		return filepath.Base(os.Args[0])
	}
	selfDir := filepath.Dir(execPath)

	// Filter our directory and subdirectories out of PATH to prevent infinite recursion
	var filtered []string
	for _, p := range strings.Split(os.Getenv("PATH"), ":") {
		if !strings.HasPrefix(p, selfDir) {
			filtered = append(filtered, p)
		}
	}
	os.Setenv("PATH", strings.Join(filtered, ":"))

	return filepath.Base(os.Args[0])
}

// ShouldWrap returns true if CLAUDECODE is set and we're running directly from Claude.
// Fast path: if NERF_IN_SCRIPT is set, we're inside a script - don't wrap.
// Slow path: check if grandparent process is "claude".
func ShouldWrap() bool {
	if os.Getenv("CLAUDECODE") == "" {
		return false
	}
	if os.Getenv("NERF_IN_SCRIPT") != "" {
		return false
	}
	return isGrandparentClaude()
}

// isGrandparentClaude checks if the grandparent process is Claude Code.
// This detects: command → shell → claude (direct invocation)
// vs: command → bash (script) → shell → claude (inside a script)
func isGrandparentClaude() bool {
	ppid := os.Getppid()

	// Read parent's parent PID from /proc
	statPath := filepath.Join("/proc", fmt.Sprintf("%d", ppid), "stat")
	data, err := os.ReadFile(statPath)
	if err != nil {
		// Can't read /proc, fall back to assuming we should wrap
		return true
	}

	// Parse stat file to get PPID (4th field)
	// Format: pid (comm) state ppid ...
	fields := strings.Fields(string(data))
	if len(fields) < 4 {
		return true
	}

	// Find the closing paren of comm field, PPID is after that
	statStr := string(data)
	closeParenIdx := strings.LastIndex(statStr, ")")
	if closeParenIdx == -1 {
		return true
	}
	afterComm := strings.Fields(statStr[closeParenIdx+1:])
	if len(afterComm) < 2 {
		return true
	}
	grandppidStr := afterComm[1] // state is [0], ppid is [1]

	var grandppid int
	fmt.Sscanf(grandppidStr, "%d", &grandppid)

	// Check if grandparent is claude
	commPath := filepath.Join("/proc", fmt.Sprintf("%d", grandppid), "comm")
	comm, err := os.ReadFile(commPath)
	if err != nil {
		return true
	}

	name := strings.TrimSpace(string(comm))
	return name == "claude"
}

// ExecReal finds and execs the real command, never returns
func ExecReal(cmd string, args []string) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		os.Stderr.WriteString(cmd + ": command not found\n")
		os.Exit(127)
	}
	argv := append([]string{cmd}, args...)
	syscall.Exec(path, argv, os.Environ())
}

// Exec execs the given path, never returns
func Exec(path string, args []string) {
	argv := append([]string{path}, args...)
	err := syscall.Exec(path, argv, os.Environ())
	if err != nil {
		os.Stderr.WriteString("exec failed: " + err.Error() + "\n")
		os.Exit(1)
	}
}

// HasFileArgs returns true if any argument is an existing file
func HasFileArgs(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		if _, err := os.Stat(arg); err == nil {
			return true
		}
	}
	return false
}
