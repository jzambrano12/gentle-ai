package rtk

import (
	"fmt"
	"os/exec"
	"strings"
)

var (
	lookPath    = exec.LookPath
	execCommand = exec.Command
)

// VerifyInstalled checks whether the rtk binary is available on PATH.
func VerifyInstalled() error {
	if _, err := lookPath("rtk"); err != nil {
		return fmt.Errorf("rtk binary not found in PATH: %w", err)
	}
	return nil
}

// VerifyVersion runs "rtk --version" and returns the trimmed output.
func VerifyVersion() (string, error) {
	cmd := execCommand("rtk", "--version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("rtk version command failed: %w", err)
	}

	version := strings.TrimSpace(string(out))
	if version == "" {
		return "", fmt.Errorf("rtk version returned empty output")
	}

	return version, nil
}
