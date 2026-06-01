//go:build linux

package procenv

import (
	"bytes"
	"os"
	"strconv"
	"strings"
)

// readProcEnv reads one environment variable from pid's /proc environ.
func readProcEnv(pid int, key string) (string, bool) {
	data, err := os.ReadFile("/proc/" + strconv.Itoa(pid) + "/environ")
	if err != nil {
		return "", false
	}
	return ParseEnviron(data, key)
}

// ParseEnviron extracts key from a Linux /proc/<pid>/environ blob.
func ParseEnviron(data []byte, key string) (string, bool) {
	if key == "" {
		return "", false
	}
	prefix := key + "="
	for _, part := range bytes.Split(data, []byte{0}) {
		line := string(part)
		if strings.HasPrefix(line, prefix) {
			return strings.TrimPrefix(line, prefix), true
		}
	}
	return "", false
}
