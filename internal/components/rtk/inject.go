package rtk

import (
	"github.com/gentleman-programming/gentle-ai/internal/model"
)

// AgentInitArgs returns the rtk init arguments for the given agent.
// All commands include --auto-patch so rtk patches settings files
// without interactive prompts (non-interactive mode defaults to skipping).
// The second return value is false if the agent is not supported by RTK.
func AgentInitArgs(agent model.AgentID) ([]string, bool) {
	switch agent {
	case model.AgentClaudeCode, model.AgentVSCodeCopilot:
		return []string{"init", "-g", "--auto-patch"}, true
	case model.AgentCursor:
		return []string{"init", "--agent", "cursor", "--auto-patch"}, true
	case model.AgentGeminiCLI:
		return []string{"init", "--gemini", "--auto-patch"}, true
	case model.AgentCodex:
		return []string{"init", "--codex", "--auto-patch"}, true
	default:
		// OpenCode plugin support is not yet mature — skip.
		return nil, false
	}
}

// IsGlobalInit reports whether the given args represent a global init
// (shared by Claude Code and VS Code Copilot).
func IsGlobalInit(args []string) bool {
	for _, arg := range args {
		if arg == "-g" {
			return true
		}
	}
	return false
}
