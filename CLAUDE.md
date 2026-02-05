# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Claude Code status line hook that displays custom terminal output when Claude Code starts or updates. It reads JSON input from stdin (provided by Claude Code's hook system) and outputs a formatted status line with ANSI color codes and terminal control sequences.

The output includes:
- User name and current directory
- Git branch and dirty status (with `*` suffix)
- Model display name
- Context window usage percentage and total cost
- Input/output token counts
- Context window size

## Build Commands

Build the binary:
```bash
go build -o bin/claude-code-status-line main.go
```

The binary can be referenced in Claude Code's hooks configuration (typically `~/.claude/hooks.json`) to customize the status line display.

## Architecture

### Entry Point (main.go)
- Reads JSON from stdin containing Claude Code session metadata
- Accepts optional username argument (defaults to system username if not provided)
- Unmarshals input into `model.Input` struct
- Calls encoder functions to format and colorize output elements
- Outputs a single line of formatted text to stdout

### Data Model (internal/model/input.go)
Defines the JSON structure received from Claude Code hooks:
- `Input`: Root structure containing session metadata
- `Model`: Model ID and display name
- `Workspace`: Current and project directories
- `Cost`: Total cost, duration, and lines changed
- `ContextWindow`: Token usage and context window metrics

### Encoder Package (internal/encoder/)
- `encoder.go`: `Encode()` function joins formatted strings with spaces
- `color.go`: ANSI color code functions (Gray, Orange, LightBlue, Yellow, etc.) using RGB escape sequences `\033[38;2;R;G;Bm`
- `git.go`: Git status functions that execute `git branch --show-current` and `git status --porcelain`
- `whoami.go`: `GetUserName()` retrieves system username
- `hyperlink.go`: `ClickableLink()` creates terminal hyperlinks using OSC 8 sequences (currently unused in main.go)

## Key Design Patterns

The codebase follows a simple pipeline architecture:
1. Input parsing (JSON unmarshal)
2. Data retrieval (git status, username)
3. Formatting (color encoding)
4. Output (single formatted line)

Each encoder color function wraps text in ANSI escape sequences and returns immediately - no state is maintained. Git operations are executed synchronously via `os/exec`.
