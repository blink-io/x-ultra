package i18n

import (
	"net/http"

	"github.com/blink-io/x/i18n"
)

func Default() func(http.Handler) http.Handler {
	return New(&Options{
		Resolvers: []Resolver{
			NewAcceptLanguageResolver(),
			NewQueryParamsResolver("lang"),
		},
	})
}

func New(o *Options) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, rv := range o.Resolvers {
				if lang := rv.Resolve(r); len(lang) > 0 {
					r = r.WithContext(i18n.NewContext(r.Context(), lang))
					break
				}
			}
			h.ServeHTTP(w, r)
		})
	}
}
