# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Test

```bash
make                    # Build for current platform
make linux-amd64        # Build for specific platform
make test               # Run tests in Docker (uses Alpine container)
```

To run a single test file locally:
```bash
export PATH="$PWD/bin/$(go env GOOS)-$(go env GOARCH)/commands:$PATH"
export BATS_LIB_PATH="$(brew --prefix)/lib/bats:$PWD/test"
CLAUDECODE=1 bats src/commands/cat.bats
```

## Project Purpose

Command wrappers that intercept POSIX commands when prepended to PATH. When `CLAUDECODE` env var is set, wrappers modify behavior to be safer/quieter. Without `CLAUDECODE`, commands pass through to real implementations.

## Architecture

**Single binary with symlinks**: All commands compile into one `nerf` binary. Symlinks in `bin/<platform>/commands/` point to it. The binary dispatches based on `argv[0]`.

**Auto-registration**: Each command in `src/commands/*.go` uses `init()` to register itself with `common.Register(name, handler)`. The registry lives in `src/common/registry.go`.

**PATH filtering**: `common.Init()` removes the wrapper directory from PATH before looking up real commands, preventing infinite recursion.

**Command handlers** (`src/commands/*.go`):
- `cat`: Truncates to 50 lines when stdout is terminal
- `docker`: Adds quiet flags; blocks `compose restart`
- `docker-compose`: Blocks `restart` (suggests `up -d`)
- `fd`/`fdfind`: Blocks exec flags (`-x`, `-X`, `--exec`, `--exec-batch`)
- `find`: Blocks `-delete`, `-exec`; remaps to `fd`
- `gh`: Blocks `run watch` (suggests `gh wait-ci`)
- `grep`/`head`/`tail`: Buffers piped stdin with 5s timeout
- `sed`: Blocks in-place editing
- `sleep`: Always sleeps max 1 second
- `touch`: Only allows creating new files

**Shared utilities** (`src/common/`):
- `common.go`: `Init()`, `ShouldWrap()`, `ExecReal()`, `Exec()`
- `registry.go`: Command handler registration
- `buffered_filter.go`: Timeout-based stdin buffering for grep/head/tail
