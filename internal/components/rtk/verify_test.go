package rtk

import (
	"errors"
	"os/exec"
	"testing"
)

func TestVerifyInstalled(t *testing.T) {
	original := lookPath
	t.Cleanup(func() { lookPath = original })

	lookPath = func(string) (string, error) { return "/usr/local/bin/rtk", nil }
	if err := VerifyInstalled(); err != nil {
		t.Fatalf("VerifyInstalled() error = %v", err)
	}

	lookPath = func(string) (string, error) { return "", errors.New("missing") }
	if err := VerifyInstalled(); err == nil {
		t.Fatalf("VerifyInstalled() expected missing binary error")
	}
}

func TestVerifyVersion(t *testing.T) {
	originalCmd := execCommand
	t.Cleanup(func() { execCommand = originalCmd })

	execCommand = func(name string, args ...string) *exec.Cmd {
		return exec.Command("echo", "rtk 0.28.2")
	}
	version, err := VerifyVersion()
	if err != nil {
		t.Fatalf("VerifyVersion() error = %v", err)
	}
	if version != "rtk 0.28.2" {
		t.Fatalf("VerifyVersion() = %q, want %q", version, "rtk 0.28.2")
	}
}
