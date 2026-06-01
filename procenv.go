package procenv

// ProcEnv reads one environment variable from another process by pid.
// Returns ("", false) when pid or key is invalid, the process is unavailable,
// or the key is not present.
func ProcEnv(pid int, key string) (string, bool) {
	if pid <= 0 || key == "" {
		return "", false
	}
	return readProcEnv(pid, key)
}
