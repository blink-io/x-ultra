package pprof

import (
	"net/http"
	"net/http/pprof"
)

func NewHandler(prefix string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(prefix+"/", pprof.Index)
	mux.HandleFunc(prefix+"/cmdline", pprof.Cmdline)
	mux.HandleFunc(prefix+"/profile", pprof.Profile)
	mux.HandleFunc(prefix+"/symbol", pprof.Symbol)
	mux.HandleFunc(prefix+"/trace", pprof.Trace)
	mux.HandleFunc(prefix+"/allocs", pprof.Handler("allocs").ServeHTTP)
	mux.HandleFunc(prefix+"/block", pprof.Handler("block").ServeHTTP)
	mux.HandleFunc(prefix+"/goroutine", pprof.Handler("goroutine").ServeHTTP)
	mux.HandleFunc(prefix+"/heap", pprof.Handler("heap").ServeHTTP)
	mux.HandleFunc(prefix+"/mutex", pprof.Handler("mutex").ServeHTTP)
	mux.HandleFunc(prefix+"/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
	return mux
}
