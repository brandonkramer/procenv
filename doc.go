// Package procenv reads one environment variable from another process by pid.
//
// Use [ProcEnv] when verifying child process identity from a run-id env marker.
// Linux reads /proc/<pid>/environ; macOS uses ps eww; other platforms return
// ("", false).
//
// Example:
//
//	value, ok := procenv.ProcEnv(pid, "MYAPP_RUN_ID")
//	if ok && value == runID {
//	    // pid matches expected run
//	}
package procenv
