package ipify

import (
	"runtime"
	"strings"
	"testing"
)

func Test_UserAgent(t *testing.T) {
	if !strings.Contains(strings.ToLower(USER_AGENT), runtime.GOOS) {
		t.Error("USER_AGENT doesn't contain the OS type:", USER_AGENT)
	}
	if !strings.Contains(USER_AGENT, VERSION) {
		t.Error("USER_AGENT doesn't contain the library version:", USER_AGENT)
	}
	t.Log("User agent: " + USER_AGENT)
}
