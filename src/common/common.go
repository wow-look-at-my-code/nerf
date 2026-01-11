package common

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// originalPATH stores the PATH before filtering, so child processes can use wrappers
var originalPATH string

// Init filters our directory out of PATH for LookPath and returns the command name from argv[0]
func Init() string {
	// Save original PATH for child processes
	originalPATH = os.Getenv("PATH")

	// Get the directory containing this binary
	execPath, err := os.Executable()
	if err != nil {
		return filepath.Base(os.Args[0])
	}
	selfDir := filepath.Dir(execPath)

	// Filter our directory and subdirectories out of PATH to prevent infinite recursion
	// This filtered PATH is only used for LookPath, not passed to children
	var filtered []string
	for _, p := range strings.Split(originalPATH, ":") {
		if !strings.HasPrefix(p, selfDir) {
			filtered = append(filtered, p)
		}
	}
	os.Setenv("PATH", strings.Join(filtered, ":"))

	return filepath.Base(os.Args[0])
}

// environ returns the environment with original PATH restored for child processes
func environ() []string {
	env := os.Environ()
	for i, e := range env {
		if strings.HasPrefix(e, "PATH=") {
			env[i] = "PATH=" + originalPATH
			return env
		}
	}
	// PATH not found, add it
	return append(env, "PATH="+originalPATH)
}

// ShouldWrap returns true if CLAUDECODE is set
func ShouldWrap() bool {
	return os.Getenv("CLAUDECODE") != ""
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
