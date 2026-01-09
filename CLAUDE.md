# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Test

```bash
make        # Build all binaries
make test   # Run bats tests (requires bats-core, bats-support, bats-assert)
```

To build a single binary: `go build -o <cmd> ./src/<cmd>.go`

Tests require `CLAUDECODE=1` environment variable to activate wrapper behavior.

## Project Purpose

This project provides wrapper binaries that intercept POSIX commands when prepended to PATH. When `CLAUDECODE` env var is set, these wrappers modify command behavior to be safer/quieter for Claude Code usage. Without `CLAUDECODE`, commands pass through to the real implementations.

## Architecture

**Activation mechanism**: The wrapper directory (`/mnt/ssdpool/projects/path_prefix`) is prepended to PATH. Each wrapper calls `common.Init()` which filters this path out of PATH before looking up the real command, preventing infinite recursion.

**Command wrappers** (`src/<cmd>.go`): Each command has its own Go binary:
- `cat`: Truncates to 50 lines/1000 tokens when stdout is terminal
- `sed`: Blocks in-place editing, only allows piped input
- `find`: Blocks dangerous flags (-delete, -exec), remaps to `fd`
- `docker`: Adds `--quiet-pull`, `--quiet`, `--progress=quiet` flags
- `timeout_filter`: Buffers piped stdin for 5s; if complete, runs filter; if timeout, dumps and passes through

**Symlink aliases**: Multiple commands share the same binary:
- `grep`, `head`, `tail` → symlink to `timeout_filter`
- `sleep` → symlink to `src/sleep.bash` (always sleeps 1s max)

**`src/common/common.go`**: Shared utilities:
- `Init()`: Filters PATH and returns command name from argv[0]
- `ShouldWrap()`: Returns true if `CLAUDECODE` is set
- `ExecReal()`: Finds and execs the real command
- `Exec()`: Execs a specific path
