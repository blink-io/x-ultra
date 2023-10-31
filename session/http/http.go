package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/blink-io/x/session"
)

type sm = session.Manager

// Manager holds the configuration settings for your sessions.
type Manager struct {
	*sm

	// Cookie contains the configuration settings for session cookies.
	Cookie SessionCookie

	// ErrorFunc allows you to control behavior when an error is encountered by
	// the Handle middleware. The default behavior is for HTTP 500
	// "Internal Server Error" message to be sent to the client and the error
	// logged using Go's standard logger. If a custom ErrorFunc is set, then
	// control will be passed to this instead. A typical use would be to provide
	// a function which logs the error and returns a customized HTML error page.
	ErrorFunc func(http.ResponseWriter, *http.Request, error)
}

// NewManager returns a new session manager with the default options. It is safe for
// concurrent use.
func NewManager(ops ...Option) *Manager {
	m := &Manager{
		sm:        session.NewManager(),
		ErrorFunc: defaultErrorFunc,
		Cookie: SessionCookie{
			Name:     "session",
			Domain:   "",
			HttpOnly: true,
			Path:     "/",
			Persist:  true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		},
	}
	for _, o := range ops {
		o(m)
	}
	return m
}

// Handle provides middleware which automatically loads and saves session
// data for the current request, and communicates the session token to and from
// the client in a cookie.
func (s *Manager) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Cookie")

		var token string
		cookie, err := r.Cookie(s.Cookie.Name)
		if err == nil {
			token = cookie.Value
		}

		ctx, err := s.Load(r.Context(), token)
		if err != nil {
			s.ErrorFunc(w, r, err)
			return
		}

		sr := r.WithContext(ctx)

		sw := &sessionResponseWriter{
			ResponseWriter: w,
			request:        sr,
			sessionManager: s,
		}

		next.ServeHTTP(sw, sr)

		if !sw.written {
			s.commitAndWriteSessionCookie(w, sr)
		}
	})
}

func (s *Manager) commitAndWriteSessionCookie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch s.Status(ctx) {
	case session.Modified:
		token, expiry, err := s.Commit(ctx)
		if err != nil {
			s.ErrorFunc(w, r, err)
			return
		}

		s.WriteSessionCookie(ctx, w, token, expiry)
	case session.Destroyed:
		s.WriteSessionCookie(ctx, w, "", time.Time{})
	}
}

// WriteSessionCookie writes a cookie to the HTTP response with the provided
// token as the cookie value and expiry as the cookie expiry time. The expiry
// time will be included in the cookie only if the session is set to persist
// or has had RememberMe(true) called on it. If expiry is an empty time.Time
// struct (so that it's IsZero() method returns true) the cookie will be
// marked with a historical expiry time and negative max-age (so the browser
// deletes it).
//
// Most applications will use the Handle() middleware and will not need to
// use this method.
func (s *Manager) WriteSessionCookie(ctx context.Context, w http.ResponseWriter, token string, expiry time.Time) {
	cookie := &http.Cookie{
		Name:     s.Cookie.Name,
		Value:    token,
		Path:     s.Cookie.Path,
		Domain:   s.Cookie.Domain,
		Secure:   s.Cookie.Secure,
		HttpOnly: s.Cookie.HttpOnly,
		SameSite: s.Cookie.SameSite,
	}

	if expiry.IsZero() {
		cookie.Expires = time.Unix(1, 0)
		cookie.MaxAge = -1
	} else if s.Cookie.Persist || s.GetBool(ctx, "__rememberMe") {
		cookie.Expires = time.Unix(expiry.Unix()+1, 0)        // Round up to the nearest second.
		cookie.MaxAge = int(time.Until(expiry).Seconds() + 1) // Round up to the nearest second.
	}

	w.Header().Add("Set-Cookie", cookie.String())
	w.Header().Add("Cache-Control", `no-cache="Set-Cookie"`)
}

func defaultErrorFunc(w http.ResponseWriter, r *http.Request, err error) {
	log.Output(2, err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

type sessionResponseWriter struct {
	http.ResponseWriter
	request        *http.Request
	sessionManager *Manager
	written        bool
}

func (sw *sessionResponseWriter) Write(b []byte) (int, error) {
	if !sw.written {
		sw.sessionManager.commitAndWriteSessionCookie(sw.ResponseWriter, sw.request)
		sw.written = true
	}

	return sw.ResponseWriter.Write(b)
}

func (sw *sessionResponseWriter) WriteHeader(code int) {
	if !sw.written {
		sw.sessionManager.commitAndWriteSessionCookie(sw.ResponseWriter, sw.request)
		sw.written = true
	}

	sw.ResponseWriter.WriteHeader(code)
}

func (sw *sessionResponseWriter) Unwrap() http.ResponseWriter {
	return sw.ResponseWriter
}
