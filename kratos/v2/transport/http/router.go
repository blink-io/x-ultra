package http

import (
	"net/http"
	"path"
)

type Router interface {
	Group(prefix string, filters ...FilterFunc) Router
	Handle(method, relativePath string, h HandlerFunc, filters ...FilterFunc)
	HandlePrefix(method, relativePath string, h HandlerFunc, filters ...FilterFunc)

	CONNECT(path string, h HandlerFunc, m ...FilterFunc)
	PATCH(path string, h HandlerFunc, m ...FilterFunc)
	PUT(path string, h HandlerFunc, m ...FilterFunc)
	POST(path string, h HandlerFunc, m ...FilterFunc)
	GET(path string, h HandlerFunc, m ...FilterFunc)
	HEAD(path string, h HandlerFunc, m ...FilterFunc)
	DELETE(path string, h HandlerFunc, m ...FilterFunc)
	TRACE(path string, h HandlerFunc, m ...FilterFunc)
	OPTIONS(path string, h HandlerFunc, m ...FilterFunc)

	Server() Server
	Prefix() string
	Filters() []FilterFunc
}

// Router is an HTTP router.
type router struct {
	prefix  string
	srv     Server
	filters []FilterFunc
}

var _ Router = (*router)(nil)

func NewRouter(prefix string, srv Server, filters ...FilterFunc) Router {
	r := &router{
		prefix:  prefix,
		srv:     srv,
		filters: filters,
	}
	return r
}

func (r *router) Filters() []FilterFunc {
	return r.filters
}

func (r *router) Prefix() string {
	return r.prefix
}

func (r *router) Server() Server {
	return r.srv
}

// Group returns a new router group.
func (r *router) Group(prefix string, filters ...FilterFunc) Router {
	var newFilters []FilterFunc
	newFilters = append(newFilters, r.filters...)
	newFilters = append(newFilters, filters...)
	return NewRouter(path.Join(r.prefix, prefix), r.srv, newFilters...)
}

// Handle registers a new route with a matcher for the URL path and method.
func (r *router) Handle(method, relativePath string, h HandlerFunc, filters ...FilterFunc) {
	r.doHandle(false, method, relativePath, h, filters...)
}

func (r *router) HandlePrefix(method, relativePath string, h HandlerFunc, filters ...FilterFunc) {
	r.doHandle(true, method, relativePath, h, filters...)
}

func (r *router) doHandle(pathPrefix bool, method, relativePath string, h HandlerFunc, filters ...FilterFunc) {
	next := http.Handler(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := &wrapper{router: r}
		ctx.Reset(res, req)
		if err := h(ctx); err != nil {
			r.srv.EncodeError()(res, req, err)
		}
	}))
	next = FilterChain(filters...)(next)
	next = FilterChain(r.filters...)(next)
	rpath := path.Join(r.prefix, relativePath)
	if pathPrefix {
		r.srv.Router().PathPrefix(rpath).Handler(next).Methods(method)
	} else {
		r.srv.Router().Handle(rpath, next).Methods(method)
	}
}

// GET registers a new GET route for a path with matching handler in the router.
func (r *router) GET(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodGet, path, h, m...)
}

// HEAD registers a new HEAD route for a path with matching handler in the router.
func (r *router) HEAD(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodHead, path, h, m...)
}

// POST registers a new POST route for a path with matching handler in the router.
func (r *router) POST(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodPost, path, h, m...)
}

// PUT registers a new PUT route for a path with matching handler in the router.
func (r *router) PUT(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodPut, path, h, m...)
}

// PATCH registers a new PATCH route for a path with matching handler in the router.
func (r *router) PATCH(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodPatch, path, h, m...)
}

// DELETE registers a new DELETE route for a path with matching handler in the router.
func (r *router) DELETE(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodDelete, path, h, m...)
}

// CONNECT registers a new CONNECT route for a path with matching handler in the router.
func (r *router) CONNECT(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodConnect, path, h, m...)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the router.
func (r *router) OPTIONS(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodOptions, path, h, m...)
}

// TRACE registers a new TRACE route for a path with matching handler in the router.
func (r *router) TRACE(path string, h HandlerFunc, m ...FilterFunc) {
	r.Handle(http.MethodTrace, path, h, m...)
}
