package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/blink-io/x/session"
	"github.com/blink-io/x/session/http/resolver"
	"github.com/blink-io/x/session/http/resolver/cookie"
)

type sm = session.Manager

// Manager holds the configuration settings for your sessions.
type Manager struct {
	*sm

	rv resolver.Resolver

	// ErrorFunc allows you to control behavior when an error is encountered by
	// the Handle middleware. The default behavior is for HTTP 500
	// "Internal Server Error" message to be sent to the client and the error
	// logged using Go's standard logger. If a custom ErrorFunc is set, then
	// control will be passed to this instead. A typical use would be to provide
	// a function which logs the error and returns a customized HTML error page.
	errorFunc func(http.ResponseWriter, *http.Request, error)
}

// NewManager returns a new session manager with the default options. It is safe for
// concurrent use.
func NewManager(ops ...Option) *Manager {
	m := &Manager{
		sm:        session.NewManager(),
		rv:        cookie.Default(),
		errorFunc: defaultErrorFunc,
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
		if s.rv == nil {
			s.ErrorFunc(w, r, errors.New("http session resolver is required"))
		} else {
			err := s.rv.Resolve(s, w, r, next)
			if err != nil {
				s.ErrorFunc(w, r, err)
			}
		}
	})
}
func (s *Manager) ErrorFunc(w http.ResponseWriter, r *http.Request, err error) {
	s.errorFunc(w, r, err)
}

func defaultErrorFunc(w http.ResponseWriter, r *http.Request, err error) {
	log.Output(2, err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
