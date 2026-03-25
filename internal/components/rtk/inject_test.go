package rtk

import (
	"testing"

	"github.com/gentleman-programming/gentle-ai/internal/model"
)

func TestAgentInitArgs(t *testing.T) {
	tests := []struct {
		agent    model.AgentID
		wantArgs []string
		wantOK   bool
	}{
		{model.AgentClaudeCode, []string{"init", "-g", "--auto-patch"}, true},
		{model.AgentVSCodeCopilot, []string{"init", "-g", "--auto-patch"}, true},
		{model.AgentCursor, []string{"init", "--agent", "cursor", "--auto-patch"}, true},
		{model.AgentGeminiCLI, []string{"init", "--gemini", "--auto-patch"}, true},
		{model.AgentCodex, []string{"init", "--codex", "--auto-patch"}, true},
		{model.AgentOpenCode, nil, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.agent), func(t *testing.T) {
			args, ok := AgentInitArgs(tt.agent)
			if ok != tt.wantOK {
				t.Fatalf("AgentInitArgs(%q) ok = %v, want %v", tt.agent, ok, tt.wantOK)
			}
			if len(args) != len(tt.wantArgs) {
				t.Fatalf("AgentInitArgs(%q) = %v, want %v", tt.agent, args, tt.wantArgs)
			}
			for i, arg := range args {
				if arg != tt.wantArgs[i] {
					t.Fatalf("AgentInitArgs(%q)[%d] = %q, want %q", tt.agent, i, arg, tt.wantArgs[i])
				}
			}
		})
	}
}

func TestIsGlobalInit(t *testing.T) {
	if !IsGlobalInit([]string{"init", "-g", "--auto-patch"}) {
		t.Fatal("IsGlobalInit should return true for global init args")
	}
	if IsGlobalInit([]string{"init", "--gemini", "--auto-patch"}) {
		t.Fatal("IsGlobalInit should return false for non-global init args")
	}
	if IsGlobalInit([]string{"init"}) {
		t.Fatal("IsGlobalInit should return false for incomplete args")
	}
}
