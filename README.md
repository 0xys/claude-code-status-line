# Claude Code Status Line

A custom status line hook for [Claude Code](https://claude.ai/code) that displays rich session information with colorized output in your terminal.

## Features

Displays the following information when Claude Code starts or updates:
- **User and directory**: Current username and working directory
- **Git status**: Current branch with dirty indicator (`*`)
- **Model info**: Active Claude model name
- **Usage metrics**: Context window usage percentage and total cost
- **Token counts**: Input and output tokens
- **Context size**: Total context window size

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

## Example Output

```
<my_username>:~/code/personal/project main* 👾 Claude Sonnet 4.5 used 45.23% $1.25 ↓15234 ↑8912 ctx 200000
```

- Gray: username and directory
- Orange: git branch and status
- Light Blue: model name
- Yellow: usage percentage and cost
- Light Red: token counts
- Magenta: context window size

## Requirements

- Go 1.25.6 or later
- Git (for git status detection)
- Claude Code CLI

## Development

See [CLAUDE.md](./CLAUDE.md) for architecture details and development guidance.

## License

MIT
