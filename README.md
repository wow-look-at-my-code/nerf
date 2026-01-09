# nerf

Command wrappers that make POSIX commands safer for Claude Code usage.

## How it works

Prepend the `bin/<platform>/commands` directory to your PATH. When `CLAUDECODE` environment variable is set, the wrappers modify command behavior to be safer/quieter. Without `CLAUDECODE`, commands pass through to the real implementations.

## Installation

```bash
# Build for your platform
make

# Add to PATH (example for macOS ARM)
export PATH="/path/to/nerf/bin/darwin-arm64/commands:$PATH"

# Enable wrapper behavior
export CLAUDECODE=1
```

## Command Wrappers

| Command | Behavior |
|---------|----------|
| `cat` | Truncates output to 50 lines when stdout is a terminal |
| `docker` | Adds `--quiet-pull`, `--quiet`, `--progress=quiet` flags; blocks `compose restart` |
| `docker-compose` | Blocks `restart` (suggests `up -d` instead) |
| `fd` / `fdfind` | Blocks `-x`, `-X`, `--exec`, `--exec-batch` flags |
| `find` | Blocks `-delete`, `-exec`; remaps to `fd` when possible |
| `gh` | Blocks `run watch` (suggests `gh wait-ci` instead) |
| `grep` | Buffers piped input with 5s timeout |
| `head` | Buffers piped input with 5s timeout |
| `sed` | Blocks in-place editing; only allows piped input |
| `sleep` | Always sleeps max 1 second |
| `tail` | Buffers piped input with 5s timeout |
| `touch` | Only allows creating new files; blocks modifying existing |

## Building

```bash
make                    # Build for current platform
make linux-amd64        # Build for Linux AMD64
make darwin-arm64       # Build for macOS ARM64
make test               # Run tests (requires Docker)
```

## License

MIT
