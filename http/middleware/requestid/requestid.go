package requestid

import (
	"net/http"

	"github.com/blink-io/x/requestid"
)

type Options = requestid.Options

var DefaultOptions = requestid.DefaultOptions

func Default() func(http.Handler) http.Handler {
	return New(DefaultOptions)
}

func New(o *Options) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rid := r.Header.Get(o.Header)
			if len(rid) == 0 {
				rid = o.Generator()
			}
			r = r.WithContext(requestid.NewContext(r.Context(), rid))
			w.Header().Set(o.Header, rid)
			h.ServeHTTP(w, r)
		})
	}
}
