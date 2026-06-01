//go:build linux

package procenv_test

import (
	"os"
	"testing"

	"github.com/brandonkramer/procenv"
)

func TestProcEnvCurrentProcess(t *testing.T) {
	const key = "PROCENV_LINUX_TEST_KEY"
	t.Setenv(key, "ok")
	v, ok := procenv.ProcEnv(os.Getpid(), key)
	if !ok || v != "ok" {
		t.Fatalf("value=%q ok=%v", v, ok)
	}
}
