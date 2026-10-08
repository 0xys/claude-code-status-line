# Claude Code Status Line

A custom status line hook for [Claude Code](https://claude.ai/code) that displays rich session information with colorized output in your terminal.

## Features

Displays the following information when Claude Code starts or updates:
- **User and directory**: Current username and working directory
- **Git status**: Current branch with dirty indicator (`*`)
- **Model info**: Active Claude model name
- **Usage metrics**: Context window usage percentage and total cost
- **Plan usage**: 5-hour and 7-day rate limit usage with the local time each window resets (e.g. `5h 24% until 12:50 7d 55% until 10/09 12:50`)

## Installation

### Option 1: Using go install

```bash
go install github.com/0xys/claude-code-status-line@latest
```

The binary will be installed to your `$GOPATH/bin` directory (typically `~/go/bin`).

### Option 2: Build from source

```bash
git clone https://github.com/0xys/claude-code-status-line.git
cd claude-code-status-line
go build -o bin/claude-code-status-line main.go
```

## Usage

Configure Claude Code to use this command by editing `~/.claude/settings.json`:

```json
{
  "statusLine": {
    "type": "command",
    "command": "claude-code-status-line <my_username>"
  }
}
```

* `<my_username>` is an optional argument. If no argument is provided, it uses your **system username**.



- Gray: username and directory
- Orange: git branch and status
- Light Blue: model name
- Yellow: usage percentage and cost
- Cyan: plan usage and reset time

## Plan usage

Claude Code passes `rate_limits` to the status line only for claude.ai Pro and Max subscribers
(or behind a Claude apps gateway with spend limits), and only after the first API response in the
session. When it is absent, the plan usage segment is omitted.

Each time `rate_limits` is present, the latest values are also written to `~/.claude/usage.json`,
so other tools can read your usage and reset times without an OAuth token:

```json
{
  "updated_at": 1791338071,
  "session_id": "abc123-def456-ghi789",
  "rate_limits": {
    "five_hour": { "used_percentage": 23.5, "resets_at": 1791346080 },
    "seven_day": { "used_percentage": 55, "resets_at": 1791612270 }
  }
}
```

`updated_at` and `resets_at` are Unix epoch seconds. Set `CLAUDE_STATUS_LINE_USAGE_FILE` to write
the file elsewhere. The file is not updated while `rate_limits` is absent, so it keeps the last
known values.

## Requirements

- Go 1.24 or later
- Git (for git status detection)
- Claude Code CLI

## Development

### Build
```bash
$ mise run build
```

### Display Sample Output

```bash
$ mise run sample
<my_username>:~/code/personal/project main* 👾 Claude Sonnet 4.5 used 45.23% $1.25 8912 ➜]..[➜ 15234
```

See [CLAUDE.md](./CLAUDE.md) for architecture details and development guidance.

## License

MIT
