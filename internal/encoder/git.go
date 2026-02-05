package encoder

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetGitStatus() (string, error) {
	branch, err := getBranchName()
	if err != nil {
		return "", err
	}

	dirty, err := isGitDirty()
	if err != nil {
		return "", err
	}
	if dirty {
		return fmt.Sprintf("%s*", branch), nil
	}
	return fmt.Sprintf("%s", branch), nil
}

func getBranchName() (string, error) {
	branch, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(branch), "\n", ""), nil
}

func isGitDirty() (bool, error) {
	statusMessage, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return false, err
	}
	return string(statusMessage) != "", nil
}
