package cookie

import (
	"context"
	"net/http"
	"time"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver"
	. "github.com/blink-io/x/session/http/shared"
)

// SessionCookie contains the configuration settings for session cookies.
type SessionCookie struct {
	// Name sets the name of the session cookie. It should not contain
	// whitespace, commas, colons, semicolons, backslashes, the equals sign or
	// control characters as per RFC6265. The default cookie name is "session".
	// If your application uses two different sessions, you must make sure that
	// the cookie name for each is unique.
	Name string

	// Domain sets the 'Domain' attribute on the session cookie. By default
	// it will be set to the domain name that the cookie was issued from.
	Domain string

	// HttpOnly sets the 'HttpOnly' attribute on the session cookie. The
	// default value is true.
	HttpOnly bool

	// Path sets the 'Path' attribute on the session cookie. The default value
	// is "/". Passing the empty string "" will result in it being set to the
	// path that the cookie was issued from.
	Path string

	// Persist sets whether the session cookie should be persistent or not
	// (i.e. whether it should be retained after a user closes their browser).
	// The default value is true, which means that the session cookie will not
	// be destroyed when the user closes their browser and the appropriate
	// 'Expires' and 'MaxAge' values will be added to the session cookie. If you
	// want to only persist some sessions (rather than all of them), then set this
	// to false and call the SetRememberMe() method for the specific sessions that you
	// want to persist.
	Persist bool

	// SameSite controls the value of the 'SameSite' attribute on the session
	// cookie. By default, this is set to 'SameSite=Lax'. If you want no SameSite
	// attribute or value in the session cookie then you should set this to 0.
	SameSite http.SameSite

	// Secure sets the 'Secure' attribute on the session cookie. The default
	// value is false. It's recommended that you set this to true and serve all
	// requests over HTTPS in production environments.
	// See https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Session_Management_Cheat_Sheet.md#transport-layer-security.
	Secure bool
}

const DefaultRememberMe = "__rememberMe"

var DefaultSessionCookie = SessionCookie{
	Name:     "X-Session-Id",
	Domain:   "",
	HttpOnly: true,
	Path:     "/",
	Persist:  true,
	Secure:   false,
	SameSite: http.SameSiteLaxMode,
}

type rv struct {
	Cookie     SessionCookie
	RememberMe string
}

func Default() resolver.Resolver {
	return New(DefaultSessionCookie)
}

func New(sc SessionCookie) resolver.Resolver {
	return &rv{
		Cookie:     sc,
		RememberMe: DefaultRememberMe,
	}
}

func (v *rv) Resolve(m resolver.Manager, ef resolver.ErrorFunc, w http.ResponseWriter, r *http.Request, next http.Handler) error {
	w.Header().Add("Vary", "Cookie")

	var token string
	cookie, err := r.Cookie(v.Cookie.Name)
	if err == nil {
		token = cookie.Value
	}

	ctx, err := m.Load(r.Context(), token)
	if err != nil {
		return err
	}

	sr := r.WithContext(ctx)

	sw := &SessionResponseWriter{
		ResponseWriter: w,
		Request:        sr,
		CommitAndWriteSession: func(w http.ResponseWriter, r *http.Request) {
			v.commitAndWriteSessionCookie(m, ef, w, sr)
		},
	}

	next.ServeHTTP(sw, sr)

	if !sw.IsWritten() {
		v.commitAndWriteSessionCookie(m, ef, w, sr)
	}
	return nil
}
func (v *rv) SessionCookie() SessionCookie {
	return v.Cookie
}

func (v *rv) commitAndWriteSessionCookie(m resolver.Manager, ef resolver.ErrorFunc, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch m.Status(ctx) {
	case session.Modified:
		token, expiry, err := m.Commit(ctx)
		if err != nil {
			ef(w, r, err)
			return
		}

		v.writeSessionCookie(ctx, m, w, token, expiry)
	case session.Destroyed:
		v.writeSessionCookie(ctx, m, w, "", time.Time{})
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
func (v *rv) writeSessionCookie(ctx context.Context, m resolver.Manager, w http.ResponseWriter, token string, expiry time.Time) {
	cookie := &http.Cookie{
		Name:     v.Cookie.Name,
		Value:    token,
		Path:     v.Cookie.Path,
		Domain:   v.Cookie.Domain,
		Secure:   v.Cookie.Secure,
		HttpOnly: v.Cookie.HttpOnly,
		SameSite: v.Cookie.SameSite,
	}

	if expiry.IsZero() {
		cookie.Expires = time.Unix(1, 0)
		cookie.MaxAge = -1
	} else if v.Cookie.Persist || m.IsRememberMe(ctx, v.RememberMe) {
		cookie.Expires = time.Unix(expiry.Unix()+1, 0)        // Round up to the nearest second.
		cookie.MaxAge = int(time.Until(expiry).Seconds() + 1) // Round up to the nearest second.
	}

	w.Header().Add("Set-Cookie", cookie.String())
	w.Header().Add("Cache-Control", `no-cache="Set-Cookie"`)
}
