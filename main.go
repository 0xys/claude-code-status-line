package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/0xys/claude-code-status-line/internal/encoder"
	"github.com/0xys/claude-code-status-line/internal/model"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	jsonInput, err := io.ReadAll(stdin)
	if err != nil {
		fmt.Printf("Failed to read stdin: %v", err)
		os.Exit(1)
	}

	var userName string
	if len(os.Args) > 1 {
		userName = os.Args[1]
	} else {
		userName, err = encoder.GetUserName()
		if err != nil {
			fmt.Printf("Failed to get user name: %v", err)
			os.Exit(1)
		}
	}

	var input model.Input
	err = json.Unmarshal([]byte(jsonInput), &input)
	if err != nil {
		fmt.Printf("Failed to unmarshal input: %v", err)
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home directory: %v", err)
		os.Exit(1)
	}
	currentDir := strings.Replace(input.Workspace.CurrentDir, home, "~", 1)

	isGitDirty, err := encoder.IsGitDirty()
	if err != nil {
		fmt.Printf("Failed to get git dirty: %v", err)
		os.Exit(1)
	}

	branchName, err := encoder.GetBranchName()
	if err != nil {
		fmt.Printf("Failed to get git branch: %v", err)
		os.Exit(1)
	}

	var gitStatus string
	if isGitDirty {
		gitStatus = encoder.Orange(fmt.Sprintf(" %s", branchName)) + encoder.DarkOrange("*")
	} else {
		gitStatus = encoder.Orange(fmt.Sprintf(" %s", branchName))
	}

	now := time.Now()

	// Persist the latest rate limits so that other tools can read them.
	// A failed write must not break the status line, so the error is ignored.
	_ = encoder.WriteRateLimits(usageFilePath(home), input, now)

	elements := []string{
		// user name and current directory
		encoder.Gray(fmt.Sprintf("%s:%s", userName, currentDir)),

		// git status
		gitStatus,

		// model display name
		encoder.LightBlue(fmt.Sprintf("👾 %s", input.Model.DisplayName)),

		// used percentage and total cost
		encoder.Yellow(fmt.Sprintf("used %.2f%% $%.2f", input.ContextWindow.UsedPercentage, input.Cost.TotalCostUsd)),
	}

	// plan usage and the next reset time (absent before the first API response)
	if rateLimits := encoder.RateLimits(input.RateLimits, now); rateLimits != "" {
		elements = append(elements, rateLimits)
	}

	fmt.Println(encoder.Encode(elements...))
}

// usageFilePath returns $CLAUDE_STATUS_LINE_USAGE_FILE, or ~/.claude/usage.json.
func usageFilePath(home string) string {
	if path := os.Getenv("CLAUDE_STATUS_LINE_USAGE_FILE"); path != "" {
		return path
	}
	return filepath.Join(home, ".claude", "usage.json")
}
