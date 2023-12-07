package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver"
	"github.com/blink-io/x/session/http/resolver/cookie"
)

// SessionHandler holds the configuration settings for your sessions.
type SessionHandler struct {
	manager session.Manager

	resolver resolver.Resolver

	// errFunc allows you to control behavior when an error is encountered by
	// the Handle middleware. The default behavior is for HTTP 500
	// "Internal Server Error" message to be sent to the client and the error
	// logged using Go's standard logger. If a custom ErrorFunc is set, then
	// control will be passed to this instead. A typical use would be to provide
	// a function which logs the error and returns a customized HTML error page.
	errFunc func(http.ResponseWriter, *http.Request, error)
}

// NewSessionHandler returns a new session manager with the default options. It is safe for
// concurrent use.
func NewSessionHandler(ops ...Option) *SessionHandler {
	m := &SessionHandler{
		manager:  session.NewManager(),
		resolver: cookie.Default(),
		errFunc:  defaultErrorFunc,
	}
	for _, o := range ops {
		o(m)
	}
	return m
}

func Handle(sh *SessionHandler) func(next http.Handler) http.Handler {
	return sh.Handle
}

// Handle provides middleware which automatically loads and saves session
// data for the current request, and communicates the session token to and from
// the client in a cookie.
func (sh *SessionHandler) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sh.resolver == nil {
			sh.errFunc(w, r, errors.New("http session resolver is required"))
		} else {
			sm := sh.manager
			nctx := session.NewContext(r.Context(), sm)
			nr := r.WithContext(nctx)
			err := sh.resolver.Resolve(sm, sh.errFunc, w, nr, next)
			if err != nil {
				sh.errFunc(w, r, err)
			}
		}
	})
}

func (sh *SessionHandler) SessionManager() session.Manager {
	return sh.manager
}

func (sh *SessionHandler) SetDefaults() {
	if sh == nil {
		return
	}
	if sh.resolver == nil {
		sh.resolver = cookie.Default()
	}
	if sh.manager == nil {
		sh.manager = session.NewManager()
	}
	if sh.errFunc == nil {
		sh.errFunc = defaultErrorFunc
	}
}

func defaultErrorFunc(w http.ResponseWriter, r *http.Request, err error) {
	log.Output(2, err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
