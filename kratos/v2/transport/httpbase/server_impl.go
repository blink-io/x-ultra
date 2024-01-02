package httpbase

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/blink-io/x/kratos/v2/internal/matcher"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/gorilla/mux"
)

var (
	_ transport.Server     = (*server)(nil)
	_ transport.Endpointer = (*server)(nil)
	_ http.Handler         = (*server)(nil)
	_ Server               = (*server)(nil)
	_ Validator            = (*server)(nil)
)

// ServerOption is an HTTP server option.
type ServerOption func(*server)

// Network with server network.
func Network(network string) ServerOption {
	return func(s *server) {
		s.network = network
	}
}

// Address with server address.
func Address(addr string) ServerOption {
	return func(s *server) {
		s.address = addr
	}
}

// Endpoint with server address.
func Endpoint(endpoint *url.URL) ServerOption {
	return func(s *server) {
		s.endpoint = endpoint
	}
}

// Timeout with server timeout.
func Timeout(timeout time.Duration) ServerOption {
	return func(s *server) {
		s.timeout = timeout
	}
}

// Logger with server logger.
// Deprecated: use global logger instead.
func Logger(_ log.Logger) ServerOption {
	return func(s *server) {}
}

// Middleware with service middleware option.
func Middleware(m ...middleware.Middleware) ServerOption {
	return func(o *server) {
		o.middleware.Use(m...)
	}
}

// Filter with HTTP middleware option.
func Filter(filters ...FilterFunc) ServerOption {
	return func(o *server) {
		o.filters = filters
	}
}

// RequestVarsDecoder with request decoder.
func RequestVarsDecoder(dec DecodeRequestFunc) ServerOption {
	return func(o *server) {
		o.decVars = dec
	}
}

// RequestQueryDecoder with request decoder.
func RequestQueryDecoder(dec DecodeRequestFunc) ServerOption {
	return func(o *server) {
		o.decQuery = dec
	}
}

// RequestDecoder with request decoder.
func RequestDecoder(dec DecodeRequestFunc) ServerOption {
	return func(o *server) {
		o.decBody = dec
	}
}

// ResponseEncoder with response encoder.
func ResponseEncoder(en EncodeResponseFunc) ServerOption {
	return func(o *server) {
		o.encResp = en
	}
}

// ErrorEncoder with error encoder.
func ErrorEncoder(en EncodeErrorFunc) ServerOption {
	return func(o *server) {
		o.encErr = en
	}
}

// TLSConfig with TLS config.
func TLSConfig(c *tls.Config) ServerOption {
	return func(o *server) {
		o.tlsConf = c
	}
}

// StrictSlash is with mux's StrictSlash
// If true, when the path pattern is "/path/", accessing "/path" will
// redirect to the former and vice versa.
func StrictSlash(strictSlash bool) ServerOption {
	return func(o *server) {
		o.strictSlash = strictSlash
	}
}

// PathPrefix with mux's PathPrefix, router will be replaced by a subrouter that start with prefix.
func PathPrefix(prefix string) ServerOption {
	return func(s *server) {
		s.router = s.router.PathPrefix(prefix).Subrouter()
	}
}

func Adapter(adapter ServerAdapter) ServerOption {
	return func(s *server) {
		s.adapter = adapter
	}
}

type serverOptions struct {
	cxt         context.Context
	tlsConf     *tls.Config
	endpoint    *url.URL
	network     string
	address     string
	timeout     time.Duration
	filters     []FilterFunc
	middleware  matcher.Matcher
	decVars     DecodeRequestFunc
	decQuery    DecodeRequestFunc
	decBody     DecodeRequestFunc
	encResp     EncodeResponseFunc
	encErr      EncodeErrorFunc
	strictSlash bool
	router      *mux.Router
}

// server is an HTTP server wrapper.
type server struct {
	cxt         context.Context
	adapter     ServerAdapter
	tlsConf     *tls.Config
	endpoint    *url.URL
	network     string
	address     string
	timeout     time.Duration
	filters     []FilterFunc
	middleware  matcher.Matcher
	decVars     DecodeRequestFunc
	decQuery    DecodeRequestFunc
	decBody     DecodeRequestFunc
	encResp     EncodeResponseFunc
	encErr      EncodeErrorFunc
	strictSlash bool
	router      *mux.Router
}

// NewServer creates an HTTP server by options.
func NewServer(opts ...ServerOption) Server {
	srv := &server{
		cxt:         context.Background(),
		network:     "tcp",
		address:     ":0",
		timeout:     1 * time.Second,
		middleware:  matcher.New(),
		decVars:     DefaultRequestVars,
		decQuery:    DefaultRequestQuery,
		decBody:     DefaultRequestDecoder,
		encResp:     DefaultResponseEncoder,
		encErr:      DefaultErrorEncoder,
		strictSlash: true,
		router:      mux.NewRouter(),
	}
	for _, o := range opts {
		o(srv)
	}
	srv.router.StrictSlash(srv.strictSlash)
	srv.router.NotFoundHandler = http.DefaultServeMux
	srv.router.MethodNotAllowedHandler = http.DefaultServeMux
	srv.router.Use(srv.filter())

	if iadapter, ok := srv.adapter.(AdapterInitializer); ok {
		aopts := &AdapterOptions{
			network:  srv.network,
			address:  srv.address,
			endpoint: srv.endpoint,
			tlsConf:  srv.tlsConf,
			handler:  FilterChain(srv.filters...)(srv.router),
		}
		iadapter.Init(srv.cxt, aopts)
	}

	return srv
}

func (s *server) Validate(ctx context.Context) error {
	if s.adapter == nil {
		return errors.New("server: adapter is required")
	}
	if v, ok := (s.adapter).(Validator); ok {
		if err := v.Validate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) DecodeVars() DecodeRequestFunc {
	return s.decVars
}

func (s *server) DecodeQuery() DecodeRequestFunc {
	return s.decQuery
}

func (s *server) DecodeBody() DecodeRequestFunc {
	return s.decBody
}

func (s *server) EncodeResponse() EncodeResponseFunc {
	return s.encResp
}

func (s *server) EncodeError() EncodeErrorFunc {
	return s.encErr
}

func (s *server) Middleware() matcher.Matcher {
	return s.middleware
}

func (s *server) Router() *mux.Router {
	return s.router
}

// Use uses a service middleware with selector.
// selector:
//   - '/*'
//   - '/helloworld.v1.Greeter/*'
//   - '/helloworld.v1.Greeter/SayHello'
func (s *server) Use(selector string, m ...middleware.Middleware) {
	s.middleware.Add(selector, m...)
}

// WalkRoute walks the router and all its sub-routers, calling walkFn for each route in the tree.
func (s *server) WalkRoute(fn WalkRouteFunc) error {
	return s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, err := route.GetMethods()
		if err != nil {
			return nil // ignore no methods
		}
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		for _, method := range methods {
			if err := fn(RouteInfo{Method: method, Path: path}); err != nil {
				return err
			}
		}
		return nil
	})
}

// WalkHandle walks the router and all its sub-routers, calling walkFn for each route in the tree.
func (s *server) WalkHandle(handle func(method, path string, handler http.HandlerFunc)) error {
	return s.WalkRoute(func(r RouteInfo) error {
		handle(r.Method, r.Path, s.ServeHTTP)
		return nil
	})
}

// Route registers an HTTP router.
func (s *server) Route(prefix string, filters ...FilterFunc) Router {
	return NewRouter(prefix, s, filters...)
}

// Handle registers a new route with a matcher for the URL path.
func (s *server) Handle(path string, h http.Handler) {
	s.router.Handle(path, h)
}

// HandlePrefix registers a new route with a matcher for the URL path prefix.
func (s *server) HandlePrefix(prefix string, h http.Handler) {
	s.router.PathPrefix(prefix).Handler(h)
}

// HandleFunc registers a new route with a matcher for the URL path.
func (s *server) HandleFunc(path string, h http.HandlerFunc) {
	s.router.HandleFunc(path, h)
}

// HandleHeader registers a new route with a matcher for the header.
func (s *server) HandleHeader(key, val string, h http.HandlerFunc) {
	s.router.Headers(key, val).Handler(h)
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.adapter.Handler().ServeHTTP(res, req)
}

func (s *server) filter() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var (
				ctx    context.Context
				cancel context.CancelFunc
			)
			if s.timeout > 0 {
				ctx, cancel = context.WithTimeout(req.Context(), s.timeout)
			} else {
				ctx, cancel = context.WithCancel(req.Context())
			}
			defer cancel()

			pathTemplate := req.URL.Path
			if route := mux.CurrentRoute(req); route != nil {
				// /path/123 -> /path/{id}
				pathTemplate, _ = route.GetPathTemplate()
			}

			tr := &Transport{
				operation:    pathTemplate,
				pathTemplate: pathTemplate,
				reqHeader:    headerCarrier(req.Header),
				replyHeader:  headerCarrier(w.Header()),
				request:      req,
				kind:         s.adapter.Kind(),
			}
			if s.endpoint != nil {
				tr.endpoint = s.endpoint.String()
			}
			tr.request = req.WithContext(transport.NewServerContext(ctx, tr))
			next.ServeHTTP(w, tr.request)
		})
	}
}

// Endpoint return a real address to registry endpoint.
// examples:
//
//	https://127.0.0.1:8000
//	Legacy: http://127.0.0.1:8000?isSecure=false
func (s *server) Endpoint() (*url.URL, error) {
	if s.endpoint != nil {
		return s.endpoint, nil
	}
	return s.adapter.Endpoint()
}

// Start start the HTTP server.
func (s *server) Start(ctx context.Context) error {
	return s.adapter.Start(ctx)
}

// Stop stop the HTTP server.
func (s *server) Stop(ctx context.Context) error {
	return s.adapter.Stop(ctx)
}

func (s *server) Listener() Listener {
	return s.adapter.Listener()
}
