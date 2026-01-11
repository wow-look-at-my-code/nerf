package common

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"time"
)

const bufferTimeout = 5 * time.Second

// RunBufferedFilter buffers stdin with a timeout, then either runs the filter
// command on the buffered input (if complete within timeout) or dumps the buffer
// and passes through remaining stdin.
func RunBufferedFilter(cmd string) {
	// If no non-option arguments, pass through (e.g., grep with no pattern shows usage)
	hasNonOption := false
	for _, arg := range os.Args[1:] {
		if len(arg) > 0 && arg[0] != '-' {
			hasNonOption = true
			break
		}
	}
	if !hasNonOption {
		ExecReal(cmd, os.Args[1:])
		return
	}

	// If file arguments are provided, pass through directly (no buffering needed)
	if HasFileArgs(os.Args[1:]) {
		ExecReal(cmd, os.Args[1:])
		return
	}

	// If stdin is not a pipe, just run command normally
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeNamedPipe) == 0 {
		ExecReal(cmd, os.Args[1:])
		return
	}

	// Buffer stdin with timeout
	var buffer []byte
	done := make(chan bool, 1)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			b, err := reader.ReadByte()
			if err != nil {
				done <- true
				return
			}
			buffer = append(buffer, b)
		}
	}()

	select {
	case <-done:
		// Completed within timeout - apply filter
		runFilter(cmd, os.Args[1:], buffer)
	case <-time.After(bufferTimeout):
		// Timeout - dump buffer and passthrough
		os.Stdout.Write(buffer)
		io.Copy(os.Stdout, os.Stdin)
	}
}

func runFilter(cmd string, args []string, input []byte) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		os.Stderr.WriteString(cmd + ": command not found\n")
		os.Exit(127)
	}

	c := exec.Command(path, args...)
	c.Stdin = &byteReader{data: input}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}

type byteReader struct {
	data []byte
	pos  int
}

func (r *byteReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}
