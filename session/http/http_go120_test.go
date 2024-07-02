//go:build go1.20

package http

import (
	"fmt"
	"github.com/blink-io/x/session"
	hdrv "github.com/blink-io/x/session/http/resolver/header"
	"net/http"
	"testing"
	"time"
)

func TestFlusher(t *testing.T) {
	t.Parallel()

	rv := hdrv.Default()
	sm := session.NewManager(session.Lifetime(500 * time.Millisecond))
	sh := NewSessionHandler(
		WithResolver(rv),
		WithSessionManager(sm),
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/get", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := w.(http.Flusher)

		fmt.Fprint(w, ok)
	}))

	ts := newTestServer(t, sh.Handle(mux))
	defer ts.Close()

	ts.execute(t, "/put")

	_, body := ts.execute(t, "/get")
	if body != "true" {
		t.Errorf("want %q; got %q", "true", body)
	}
}

func TestHijacker(t *testing.T) {
	t.Parallel()

	rv := hdrv.Default()
	sm := session.NewManager(session.Lifetime(500 * time.Millisecond))
	sh := NewSessionHandler(
		WithResolver(rv),
		WithSessionManager(sm),
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/get", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := w.(http.Hijacker)

		fmt.Fprint(w, ok)
	}))

	ts := newTestServer(t, sh.Handle(mux))
	defer ts.Close()

	ts.execute(t, "/put")

	_, body := ts.execute(t, "/get")
	if body != "true" {
		t.Errorf("want %q; got %q", "true", body)
	}
}
