package http

import (
	"net/http"
	"testing"

	headerrv "github.com/blink-io/x/session/http/resolver/header"
)

func TestIdleTimeout_Header(t *testing.T) {
	//testIdleTimeout(t, header.Default())
}

func TestEnable_Header(t *testing.T) {
	t.Parallel()

	rv := headerrv.Default()
	sessionManager := NewManager(WithResolver(rv))

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		sessionManager.Put(r.Context(), "foo", "bar")
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		s := sessionManager.Get(r.Context(), "foo").(string)
		w.Write([]byte(s))
	})

	ts := newTestServer(t, sessionManager.Handle(mux))
	defer ts.Close()

	header, _ := ts.execute(t, "/put")
	token1 := header.Get(headerrv.DefaultHeader)

	header, body := ts.execute(t, "/get")
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}
	if header.Get(headerrv.DefaultHeader) != "" {
		t.Errorf("want %q; got %q", "", header.Get(headerrv.DefaultHeader))
	}

	header, _ = ts.execute(t, "/put")
	token2 := header.Get(headerrv.DefaultHeader)
	if token1 != token2 {
		t.Error("want tokens to be the same")
	}
}
