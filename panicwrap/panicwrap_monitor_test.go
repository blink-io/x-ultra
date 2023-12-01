//go:build !windows
// +build !windows

package panicwrap

import (
	"bytes"
	"strings"
	"testing"
)

func TestPanicWrap_monitor(t *testing.T) {

	stdout := new(bytes.Buffer)

	p := helperProcess("panic-monitor")
	p.Stdout = stdout
	//p.Stderr = new(bytes.Buffer)
	if err := p.Run(); err == nil || err.Error() != "exit status 2" {
		t.Fatalf("err: %s", err)
	}

	if !strings.Contains(stdout.String(), "wrapped:") {
		t.Fatalf("didn't wrap: %#v", stdout.String())
	}
}
