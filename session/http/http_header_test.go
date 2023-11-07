package http

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver/cookie"
	headerrv "github.com/blink-io/x/session/http/resolver/header"
)

func TestIdleTimeout_Header(t *testing.T) {
	//testIdleTimeout(t, header.Default())
}

func TestEnable_Header(t *testing.T) {
	t.Parallel()

	rv := headerrv.Default()
	sessionManager := NewSessionHandler(WithResolver(rv))

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		doSessionManagerPut(r, "foo", "bar")
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		s := doSessionManagerGet(r, "foo").(string)
		w.Write([]byte(s))
	})

	ts := newTestServer(t, sessionManager.Handle(mux))
	defer ts.Close()

	header, _ := ts.execute(t, "/put")
	token1 := header.Get(headerrv.DefaultHeader)

	header, body := ts.executeWithHeaders(t, "/get", map[string]string{
		headerrv.DefaultHeader: token1,
	})
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}
	if header.Get(headerrv.DefaultHeader) != "" {
		t.Errorf("want %q; got %q", "", header.Get(headerrv.DefaultHeader))
	}

	header, _ = ts.executeWithHeaders(t, "/put", map[string]string{
		headerrv.DefaultHeader: token1,
	})
	token2 := header.Get(headerrv.DefaultHeader)
	if token1 != token2 {
		t.Error("want tokens to be the same")
	}
}

func TestLifetime_Header(t *testing.T) {
	t.Parallel()

	rv := headerrv.Default()
	sm := session.NewManager(session.Lifetime(500 * time.Millisecond))
	sh := NewSessionHandler(
		WithResolver(rv),
		WithSessionManager(sm),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		sm, ok := session.FromContext(r.Context())
		if ok {
			sm.Put(r.Context(), "foo", "bar")
		}
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		v := doSessionManagerGet(r, "foo")
		if v == nil {
			http.Error(w, "foo does not exist in session", 500)
			return
		}
		w.Write([]byte(v.(string)))
	})

	ts := newTestServer(t, sh.Handle(mux))
	defer ts.Close()

	header1, _ := ts.execute(t, "/put")
	token1 := header1.Get(headerrv.DefaultHeader)

	_, body := ts.executeWithHeaders(t, "/get", map[string]string{
		headerrv.DefaultHeader: token1,
	})
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}
	time.Sleep(time.Second)

	_, body = ts.execute(t, "/get")
	if body != "foo does not exist in session\n" {
		t.Errorf("want %q; got %q", "foo does not exist in session\n", body)
	}
}

func TestRenewToken_Header(t *testing.T) {
	t.Parallel()

	rv := headerrv.Default()
	sessionManager := NewSessionHandler(WithResolver(rv))

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		doSessionManagerPut(r, "foo", "bar")

	})
	mux.HandleFunc("/renew", func(w http.ResponseWriter, r *http.Request) {
		err := doSessionManagerRenewToken(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		v := doSessionManagerGet(r, "foo")
		if v == nil {
			http.Error(w, "foo does not exist in session", 500)
			return
		}
		w.Write([]byte(v.(string)))
	})

	ts := newTestServer(t, sessionManager.Handle(mux))
	defer ts.Close()

	header, _ := ts.execute(t, "/put")
	originalToken := header.Get(headerrv.DefaultHeader)

	header2, _ := ts.executeWithHeaders(t, "/renew", map[string]string{
		headerrv.DefaultHeader: originalToken,
	})
	newToken := header2.Get(headerrv.DefaultHeader)

	if newToken == originalToken {
		t.Fatal("token has not changed")
	}

	_, body := ts.executeWithHeaders(t, "/get", map[string]string{
		headerrv.DefaultHeader: newToken,
	})
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}
}

func TestDestroy_Header(t *testing.T) {
	t.Parallel()

	rv := headerrv.Default()
	sessionManager := NewSessionHandler(WithResolver(rv))

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		doSessionManagerPut(r, "foo", "bar")
	})
	mux.HandleFunc("/destroy", func(w http.ResponseWriter, r *http.Request) {
		headerrv.Default()
		err := doSessionManagerDestroy(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		//v := sessionManager.Get(r.Context(), "foo")
		v := doSessionManagerGet(r, "foo")
		if v == nil {
			http.Error(w, "foo does not exist in session", 500)
			return
		}
		w.Write([]byte(v.(string)))
	})

	ts := newTestServer(t, sessionManager.Handle(mux))
	defer ts.Close()

	header, _ := ts.execute(t, "/put")
	token := header.Get(headerrv.DefaultHeader)

	header2, _ := ts.executeWithHeaders(t, "/destroy", map[string]string{
		headerrv.DefaultHeader: token,
	})
	token2 := header2.Get(headerrv.DefaultHeader)

	if len(token2) != 0 {
		t.Fatalf("got %s: expected empty", token2)
	}

	_, body := ts.execute(t, "/get")
	if body != "foo does not exist in session\n" {
		t.Errorf("want %q; got %q", "foo does not exist in session\n", body)
	}
}

func doSessionManagerPut(r *http.Request, key string, val any) {
	sm, ok := session.FromContext(r.Context())
	if ok {
		sm.Put(r.Context(), key, val)
	}
}

func doSessionManagerGet(r *http.Request, key string) any {
	sm, ok := session.FromContext(r.Context())
	if ok {
		return sm.Get(r.Context(), key)
	}
	return errNoSessionManager
}

var errNoSessionManager = errors.New("no http session manager")

func doSessionManagerRenewToken(r *http.Request) error {
	sm, ok := session.FromContext(r.Context())
	if ok {
		return sm.RenewToken(r.Context())
	}
	return errNoSessionManager
}

func doSessionManagerDestroy(r *http.Request) error {
	sm, ok := session.FromContext(r.Context())
	if ok {
		return sm.Destroy(r.Context())
	}
	return errNoSessionManager
}

func doSessionManagerSetRememberMe(r *http.Request, okme bool) error {
	sm, ok := session.FromContext(r.Context())
	if ok {
		sm.SetRememberMe(r.Context(), cookie.DefaultRememberMe, okme)
		return nil
	}
	return errNoSessionManager
}
