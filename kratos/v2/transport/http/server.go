package http

import (
	"net/http"

	"github.com/blink-io/x/kratos/v2/internal/matcher"
	"github.com/blink-io/x/kratos/v2/transport/http/adapter"
	"github.com/blink-io/x/kratos/v2/util"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/gorilla/mux"
)

type ServerCodec interface {
	//DecodeVars defines decoding request function
	DecodeVars() DecodeRequestFunc

	//DecodeQuery defines decoding query strings function
	DecodeQuery() DecodeRequestFunc

	//DecodeBody defines decoding request's body function
	DecodeBody() DecodeRequestFunc

	//EncodeResponse defines encoding response function
	EncodeResponse() EncodeResponseFunc

	//EncodeError defines encoding error function
	EncodeError() EncodeErrorFunc

	//Middleware defines middlewares
	Middleware() matcher.Matcher
}

type ServerRouter interface {

	// Route registers an HTTP router.
	Route(prefix string, filters ...FilterFunc) Router

	// Handle registers a new route with a matcher for the URL path.
	Handle(path string, h http.Handler)

	// HandlePrefix registers a new route with a matcher for the URL path prefix.
	HandlePrefix(prefix string, h http.Handler)

	// HandleFunc registers a new route with a matcher for the URL path.
	HandleFunc(path string, h http.HandlerFunc)

	// HandleHeader registers a new route with a matcher for the header.
	HandleHeader(key, val string, h http.HandlerFunc)

	// WalkRoute walks the router and all its sub-routers, calling walkFn for each route in the tree.
	WalkRoute(fn WalkRouteFunc) error

	// WalkHandle walks the router and all its sub-routers, calling walkFn for each route in the tree.
	WalkHandle(handle func(method, path string, handler http.HandlerFunc)) error
}

type Validator = util.Validator

type Server interface {
	ServerRouter

	ServerCodec

	transport.Server

	transport.Endpointer

	http.Handler

	Listener() adapter.Listener

	Router() *mux.Router
}
