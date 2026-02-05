package encoder

import (
	"os/exec"
	"strings"
)

func GetBranchName() (string, error) {
	branch, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(branch), "\n", ""), nil
}

func IsGitDirty() (bool, error) {
	statusMessage, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return false, err
	}
	return string(statusMessage) != "", nil
}
