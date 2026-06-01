//go:build darwin

package procenv

import (
	"context"
	"os/exec"
	"strconv"
	"strings"
)

// readProcEnv reads one environment variable from pid via ps.
func readProcEnv(pid int, key string) (string, bool) {
	out, err := exec.CommandContext(context.Background(), "ps", "eww", "-p", strconv.Itoa(pid)).CombinedOutput() //nolint:gosec // pid comes from caller process metadata
	if err != nil {
		return "", false
	}
	return ParseOutput(string(out), key)
}

// ParseOutput extracts key from ps eww output.
func ParseOutput(out, key string) (string, bool) {
	prefix := key + "="
	for _, field := range strings.Fields(out) {
		if strings.HasPrefix(field, prefix) {
			return strings.TrimPrefix(field, prefix), true
		}
	}
	return "", false
}
