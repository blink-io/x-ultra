package http

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type MuxRouter interface {
	Use(mwf ...mux.MiddlewareFunc)
	Match(req *http.Request, match *mux.RouteMatch) bool
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	Get(name string) *mux.Route
	GetRoute(name string) *mux.Route
	StrictSlash(value bool) *mux.Router
	SkipClean(value bool) *mux.Router
	UseEncodedPath() *mux.Router
	NewRoute() *mux.Route
	Name(name string) *mux.Route
	Handle(path string, handler http.Handler) *mux.Route
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
	Headers(pairs ...string) *mux.Route
	Host(tpl string) *mux.Route
	MatcherFunc(f mux.MatcherFunc) *mux.Route
	Methods(methods ...string) *mux.Route
	Path(tpl string) *mux.Route
	PathPrefix(tpl string) *mux.Route
	Queries(pairs ...string) *mux.Route
	Schemes(schemes ...string) *mux.Route
	BuildVarsFunc(f mux.BuildVarsFunc) *mux.Route
	Walk(walkFn mux.WalkFunc) error
}

type MuxRoute interface {
	SkipClean() bool
	Match(req *http.Request, match *mux.RouteMatch) bool
	GetError() error
	BuildOnly() *mux.Route
	Handler(handler http.Handler) *mux.Route
	HandlerFunc(f func(http.ResponseWriter, *http.Request)) *mux.Route
	GetHandler() http.Handler
	Name(name string) *mux.Route
	GetName() string
	Headers(pairs ...string) *mux.Route
	HeadersRegexp(pairs ...string) *mux.Route
	Host(tpl string) *mux.Route
	MatcherFunc(f mux.MatcherFunc) *mux.Route
	Methods(methods ...string) *mux.Route
	Path(tpl string) *mux.Route
	PathPrefix(tpl string) *mux.Route
	Queries(pairs ...string) *mux.Route
	Schemes(schemes ...string) *mux.Route
	BuildVarsFunc(f mux.BuildVarsFunc) *mux.Route
	Subrouter() *mux.Router
	URL(pairs ...string) (*url.URL, error)
	URLHost(pairs ...string) (*url.URL, error)
	URLPath(pairs ...string) (*url.URL, error)
	GetPathTemplate() (string, error)
	GetPathRegexp() (string, error)
	GetQueriesRegexp() ([]string, error)
	GetQueriesTemplates() ([]string, error)
	GetMethods() ([]string, error)
	GetHostTemplate() (string, error)
	GetVarNames() ([]string, error)
	//GoString() string
}

var _ MuxRouter = (*mux.Router)(nil)
var _ MuxRoute = (*mux.Route)(nil)

type muxRouter struct {
	*mux.Router
}

func newMuxRouter() *muxRouter {
	r := &muxRouter{
		Router: mux.NewRouter(),
	}
	return r
}
