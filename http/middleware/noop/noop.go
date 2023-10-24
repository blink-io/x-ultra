package noop

import (
	"net/http"
)

// NewHandler creates a noop handler
func NewHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
