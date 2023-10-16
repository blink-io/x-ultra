package i18n

import (
	stdhttp "net/http"

	"golang.org/x/text/language"
)

// Resolver gets locale info from HTTP/gRPC transport
type Resolver interface {
	Resolve(*stdhttp.Request) string
}

var _ Resolver = (*HeaderResolver)(nil)
var _ Resolver = (*CookieResolver)(nil)
var _ Resolver = (*QueryParamsResolver)(nil)
var _ Resolver = (*AcceptLanguageResolver)(nil)
var _ Resolver = (*PathResolver)(nil)

// HeaderResolver gets language from header in HTTP
type HeaderResolver struct {
	Name string
}

func (rv *HeaderResolver) Resolve(r *stdhttp.Request) string {
	val := r.Header.Get(rv.Name)
	return val
}

// CookieResolver gets locale info from HTTP cookie
type CookieResolver struct {
	Name string
}

func (rv *CookieResolver) Resolve(r *stdhttp.Request) string {
	if c, err := r.Cookie(rv.Name); err == nil {
		return c.Value
	}
	return ""
}

// QueryParamsResolver gets locale info from query string
type QueryParamsResolver struct {
	Name string
}

func NewQueryParamsResolver(name string) Resolver {
	return &QueryParamsResolver{
		Name: name,
	}
}
func (rv *QueryParamsResolver) Resolve(r *stdhttp.Request) string {
	return r.URL.Query().Get(rv.Name)
}

// AcceptLanguageResolver gets locale info from header accept-language
// example of accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7
type AcceptLanguageResolver struct {
}

func NewAcceptLanguageResolver() Resolver {
	return &AcceptLanguageResolver{}

}

func (rv *AcceptLanguageResolver) Resolve(r *stdhttp.Request) string {
	val := r.Header.Get("Accept-Language")
	if tags, _, err := language.ParseAcceptLanguage(val); err == nil {
		return tags[0].String()
	}
	return ""
}

// PathResolver gets locale info from url path expr.
// Foe example, /products/{lang} -> /products/zh-cn?product_id=XXX,
type PathResolver struct {
	Key  string
	Expr string
}

func (rv *PathResolver) Resolve(r *stdhttp.Request) string {
	return ""
}
