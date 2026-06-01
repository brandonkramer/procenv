package procenv_test

import (
	"os"
	"testing"

	"github.com/brandonkramer/procenv"
)

func TestProcEnvBadPID(t *testing.T) {
	t.Parallel()

	if _, ok := procenv.ProcEnv(-1, "X"); ok {
		t.Fatal("expected false for bad pid")
	}
	if _, ok := procenv.ProcEnv(0, "X"); ok {
		t.Fatal("expected false for pid 0")
	}
}

func TestProcEnvEmptyKey(t *testing.T) {
	t.Parallel()

	if _, ok := procenv.ProcEnv(os.Getpid(), ""); ok {
		t.Fatal("expected false for empty key")
	}
}

func TestProcEnvCurrentMissingKey(t *testing.T) {
	t.Parallel()

	if _, ok := procenv.ProcEnv(os.Getpid(), "PROCENV_DEFINITELY_MISSING_KEY"); ok {
		t.Fatal("expected missing key")
	}
}
