//go:build darwin

package procenv_test

import (
	"testing"

	"github.com/brandonkramer/procenv"
)

func TestParseOutputSuccess(t *testing.T) {
	t.Parallel()

	v, ok := procenv.ParseOutput("123 sleep AGENTD_RUN_ID=abc", "AGENTD_RUN_ID")
	if !ok || v != "abc" {
		t.Fatalf("v=%q ok=%v", v, ok)
	}
}

func TestParseOutputMissing(t *testing.T) {
	t.Parallel()

	if _, ok := procenv.ParseOutput("1 sleep", "AGENTD_RUN_ID"); ok {
		t.Fatal("expected missing")
	}
}
