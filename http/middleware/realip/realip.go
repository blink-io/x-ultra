package realip

import (
	"net/http"

	"github.com/blink-io/x/realip"
)

type Options = realip.Options

var DefaultOptions = realip.DefaultOptions

func Default() func(http.Handler) http.Handler {
	return New(DefaultOptions)
}

func New(o *Options) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if ip := o.GetFromHTTP(r); len(ip) > 0 {
				r = r.WithContext(realip.NewContext(r.Context(), ip))
			}
			h.ServeHTTP(w, r)
		})
	}
}
