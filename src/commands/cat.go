package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"path_prefix/src/common"
)

func init() {
	common.Register("cat", Cat)
}

const (
	maxLines  = 50
	maxTokens = 1000
)

func countTokens(s string) int {
	// Simple approximation: split on whitespace
	return len(strings.Fields(s))
}

func Cat() {
	// Check if stdout is a pipe (not a terminal)
	stat, _ := os.Stdout.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Output is piped, pass through to real cat
		common.ExecReal("cat", os.Args[1:])
	}

	// No file arguments means stdin - pass through
	if len(os.Args) < 2 {
		common.ExecReal("cat", os.Args[1:])
	}

	// Process each file
	for _, filename := range os.Args[1:] {
		if strings.HasPrefix(filename, "-") {
			// It's a flag, pass through to real cat
			common.ExecReal("cat", os.Args[1:])
		}

		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %s: %v\n", filename, err)
			continue
		}

		stat, err := file.Stat()
		if err != nil {
			file.Close()
			fmt.Fprintf(os.Stderr, "cat: %s: %v\n", filename, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		lines := 0
		tokens := 0
		truncated := false

		for scanner.Scan() {
			line := scanner.Text()
			lineTokens := countTokens(line)

			if lines >= maxLines || tokens+lineTokens > maxTokens {
				truncated = true
				break
			}

			fmt.Println(line)
			lines++
			tokens += lineTokens
		}

		file.Close()

		if truncated {
			remaining := stat.Size()
			fmt.Fprintf(os.Stderr, "\n... truncated after %d lines.\n", lines)
			fmt.Fprintf(os.Stderr, "Use the Read tool to read the full file (%d bytes): %s\n", remaining, filename)
		}
	}
}
