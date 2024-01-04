package i18n

import (
	stdhttp "net/http"

	"github.com/blink-io/x/kratos/v2/transport"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/text/language"
)

type Resolver interface {
	Resolve(transport.Transporter) string
}

var _ Resolver = (*HeaderResolver)(nil)
var _ Resolver = (*CookieResolver)(nil)
var _ Resolver = (*QueryParamsResolver)(nil)
var _ Resolver = (*AcceptLanguageResolver)(nil)

// HeaderResolver gets locale info from header(HTTP) or metadata(gRPC)
type HeaderResolver struct {
	Name string
}

func (r *HeaderResolver) Resolve(tr transport.Transporter) string {
	val := tr.RequestHeader().Get(r.Name)
	return val
}

// CookieResolver gets locale info from HTTP cookie
type CookieResolver struct {
	Name string
}

func (r *CookieResolver) Resolve(tr transport.Transporter) string {
	if tr.Kind() == transport.KindHTTP {
		val := tr.RequestHeader().Get("Cookie")
		hr := stdhttp.Request{
			Header: stdhttp.Header{
				"Cookie": []string{val},
			},
		}
		if c, err := hr.Cookie(r.Name); err == nil {
			return c.Value
		}
	}
	return ""
}

// QueryParamsResolver gets locale info from query string
type QueryParamsResolver struct {
	Name string
}

func (r *QueryParamsResolver) Resolve(tr transport.Transporter) string {
	if tr.Kind() == transport.KindHTTP {
		if ht, ok := tr.(*khttp.Transport); ok {
			val := ht.Request().URL.Query().Get(r.Name)
			return val
		}
	}
	return ""
}

// AcceptLanguageResolver gets locale info from header accept-language
// example of accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7
type AcceptLanguageResolver struct {
}

func (r *AcceptLanguageResolver) Resolve(tr transport.Transporter) string {
	lang := tr.RequestHeader().Get("accept-language")
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err == nil && len(tags) > 0 {
		return tags[0].String()
	}
	return ""
}
