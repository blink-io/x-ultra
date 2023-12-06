package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver"
	"github.com/blink-io/x/session/http/resolver/cookie"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testServer struct {
	*httptest.Server
}

func newCookieResolver() resolver.Resolver {
	return cookie.Default()
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

func (ts *testServer) executeWithHeaders(t *testing.T, urlPath string, headers map[string]string) (http.Header, string) {
	req, err := http.NewRequest(http.MethodGet, ts.URL+urlPath, nil)
	for h, v := range headers {
		req.Header.Set(h, v)
	}
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.Header, string(body)
}

func (ts *testServer) execute(t *testing.T, urlPath string) (http.Header, string) {
	req, err := http.NewRequest(http.MethodGet, ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.Header, string(body)
}

func extractTokenFromCookie(c string) string {
	parts := strings.Split(c, ";")
	return strings.SplitN(parts[0], "=", 2)[1]
}

func TestEnable_Cookie(t *testing.T) {
	t.Parallel()

	sessionManager := NewSessionHandler()

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
	token1 := extractTokenFromCookie(header.Get("Set-Cookie"))

	header, body := ts.execute(t, "/get")
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}
	if header.Get("Set-Cookie") != "" {
		t.Errorf("want %q; got %q", "", header.Get("Set-Cookie"))
	}

	header, _ = ts.execute(t, "/put")
	token2 := extractTokenFromCookie(header.Get("Set-Cookie"))
	if token1 != token2 {
		t.Error("want tokens to be the same")
	}
}

func TestLifetime_Cookie(t *testing.T) {
	t.Parallel()

	sm := session.NewManager(session.Lifetime(500 * time.Millisecond))
	sessionManager := NewSessionHandler(WithSessionManager(sm))

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		//sessionManager.Put(r.Context(), "foo", "bar")
		doSessionManagerPut(r, "foo", "bar")
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

	ts.execute(t, "/put")

	_, body := ts.execute(t, "/get")
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}
	time.Sleep(time.Second)

	_, body = ts.execute(t, "/get")
	if body != "foo does not exist in session\n" {
		t.Errorf("want %q; got %q", "foo does not exist in session\n", body)
	}
}

func TestIdleTimeout_Cookie(t *testing.T) {
	t.Parallel()

	sm := session.NewManager(
		session.IdleTimeout(200*time.Millisecond),
		session.Lifetime(time.Second),
		session.TokenGenerator(func() (string, error) {
			return uuid.NewString(), nil
		}),
	)

	sessionManager := NewSessionHandler(
		WithSessionManager(sm),
		WithErrorFunc(defaultErrorFunc),
		WithResolver(cookie.Default()),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		doSessionManagerPut(r, "foo", "bar")
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

	ts.execute(t, "/put")

	time.Sleep(100 * time.Millisecond)
	ts.execute(t, "/get")

	time.Sleep(150 * time.Millisecond)
	_, body := ts.execute(t, "/get")
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}

	time.Sleep(200 * time.Millisecond)
	_, body = ts.execute(t, "/get")
	if body != "foo does not exist in session\n" {
		t.Errorf("want %q; got %q", "foo does not exist in session\n", body)
	}
}

func TestDestroy_Cookie(t *testing.T) {
	t.Parallel()

	crv := newCookieResolver()
	sessionManager := NewSessionHandler(WithResolver(crv))

	rv, ok := crv.(interface {
		SessionCookie() cookie.SessionCookie
	})
	require.Equal(t, true, ok)

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		doSessionManagerPut(r, "foo", "bar")
	})
	mux.HandleFunc("/destroy", func(w http.ResponseWriter, r *http.Request) {
		err := doSessionManagerDestroy(r)
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

	ts.execute(t, "/put")
	header, _ := ts.execute(t, "/destroy")
	cookie := header.Get("Set-Cookie")

	if strings.HasPrefix(cookie, fmt.Sprintf("%s=;", rv.SessionCookie().Name)) == false {
		t.Fatalf("got %q: expected prefix %q", cookie, fmt.Sprintf("%s=;", rv.SessionCookie().Name))
	}
	if strings.Contains(cookie, "Expires=Thu, 01 Jan 1970 00:00:01 GMT") == false {
		t.Fatalf("got %q: expected to contain %q", cookie, "Expires=Thu, 01 Jan 1970 00:00:01 GMT")
	}
	if strings.Contains(cookie, "Max-Age=0") == false {
		t.Fatalf("got %q: expected to contain %q", cookie, "Max-Age=0")
	}

	_, body := ts.execute(t, "/get")
	if body != "foo does not exist in session\n" {
		t.Errorf("want %q; got %q", "foo does not exist in session\n", body)
	}
}

func TestRenewToken_Cookie(t *testing.T) {
	t.Parallel()

	sessionManager := NewSessionHandler()

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
	cookie := header.Get("Set-Cookie")
	originalToken := extractTokenFromCookie(cookie)

	header, _ = ts.execute(t, "/renew")
	cookie = header.Get("Set-Cookie")
	newToken := extractTokenFromCookie(cookie)

	if newToken == originalToken {
		t.Fatal("token has not changed")
	}

	_, body := ts.execute(t, "/get")
	if body != "bar" {
		t.Errorf("want %q; got %q", "bar", body)
	}
}

func TestRememberMe_Cookie(t *testing.T) {
	t.Parallel()

	csc := &cookie.DefaultSessionCookie
	csc.Persist = false

	crv := cookie.New(*csc)

	sh := NewSessionHandler(WithResolver(crv))

	mux := http.NewServeMux()
	mux.HandleFunc("/put-normal", func(w http.ResponseWriter, r *http.Request) {
		doSessionManagerPut(r, "foo", "bar")
	})
	mux.HandleFunc("/put-rememberMe-true", func(w http.ResponseWriter, r *http.Request) {
		if err := doSessionManagerSetRememberMe(r, true); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		doSessionManagerPut(r, "foo", "bar")
	})
	mux.HandleFunc("/put-rememberMe-false", func(w http.ResponseWriter, r *http.Request) {
		if err := doSessionManagerSetRememberMe(r, false); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		doSessionManagerPut(r, "foo", "bar")
	})

	ts := newTestServer(t, sh.Handle(mux))
	defer ts.Close()

	header, _ := ts.execute(t, "/put-normal")
	header.Get("Set-Cookie")

	if strings.Contains(header.Get("Set-Cookie"), "Max-Age=") || strings.Contains(header.Get("Set-Cookie"), "Expires=") {
		t.Errorf("want no Max-Age or Expires attributes; got %q", header.Get("Set-Cookie"))
	}

	header, _ = ts.execute(t, "/put-rememberMe-true")
	header.Get("Set-Cookie")

	if !strings.Contains(header.Get("Set-Cookie"), "Max-Age=") || !strings.Contains(header.Get("Set-Cookie"), "Expires=") {
		t.Errorf("want Max-Age and Expires attributes; got %q", header.Get("Set-Cookie"))
	}

	header, _ = ts.execute(t, "/put-rememberMe-false")
	header.Get("Set-Cookie")

	if strings.Contains(header.Get("Set-Cookie"), "Max-Age=") || strings.Contains(header.Get("Set-Cookie"), "Expires=") {
		t.Errorf("want no Max-Age or Expires attributes; got %q", header.Get("Set-Cookie"))
	}
}

func TestIterate_Cookie(t *testing.T) {
	t.Parallel()

	sh := NewSessionHandler()
	sm := sh.SessionManager()

	mux := http.NewServeMux()
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		doSessionManagerPut(r, "foo", r.URL.Query().Get("foo"))
	})

	for i := 0; i < 3; i++ {
		ts := newTestServer(t, sh.Handle(mux))
		defer ts.Close()

		ts.execute(t, "/put?foo="+strconv.Itoa(i))
	}

	results := []string{}

	err := sm.Iterate(context.Background(), func(ctx context.Context) error {
		i := sm.GetString(ctx, "foo")
		results = append(results, i)
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	sort.Strings(results)

	if !reflect.DeepEqual(results, []string{"0", "1", "2"}) {
		t.Fatalf("unexpected value: got %v", results)
	}

	err = sm.Iterate(context.Background(), func(ctx context.Context) error {
		return errors.New("expected error")
	})
	if err.Error() != "expected error" {
		t.Fatal("didn't get expected error")
	}
}
