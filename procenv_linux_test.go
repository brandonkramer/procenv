//go:build linux

package procenv_test

import (
	"context"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/brandonkramer/procenv"
)

func TestParseEnviron(t *testing.T) {
	t.Parallel()

	blob := []byte("HOME=/tmp\x00PROCENV_KEY=abc\x00")
	v, ok := procenv.ParseEnviron(blob, "PROCENV_KEY")
	if !ok || v != "abc" {
		t.Fatalf("value=%q ok=%v", v, ok)
	}
	if _, ok := procenv.ParseEnviron(blob, "MISSING"); ok {
		t.Fatal("expected missing key")
	}
	if _, ok := procenv.ParseEnviron(blob, ""); ok {
		t.Fatal("expected empty key rejection")
	}
}

func TestProcEnvSubprocess(t *testing.T) {
	t.Parallel()

	const key = "PROCENV_LINUX_SUBPROC_KEY"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sleep", "30")
	cmd.Env = append(os.Environ(), key+"=child")
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = cmd.Process.Kill() }()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		v, ok := procenv.ProcEnv(cmd.Process.Pid, key)
		if ok && v == "child" {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("expected subprocess env marker")
}
