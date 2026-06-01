//go:build !unix

package procenv

// ProcEnv is unsupported on this platform.
func ProcEnv(_ int, _ string) (string, bool) {
	return "", false
}
