//go:build !unix

package procenv

// ProcEnv is unsupported on this platform.
func readProcEnv(_ int, _ string) (string, bool) {
	return "", false
}
