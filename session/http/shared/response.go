package shared

import (
	"net/http"
)

type SessionResponseWriter struct {
	http.ResponseWriter
	Request               *http.Request
	CommitAndWriteSession func(http.ResponseWriter, *http.Request)
	written               bool
}

func (sw *SessionResponseWriter) Write(b []byte) (int, error) {
	if !sw.written {
		if sw.CommitAndWriteSession != nil {
			sw.CommitAndWriteSession(sw.ResponseWriter, sw.Request)
		}
		sw.written = true
	}

	return sw.ResponseWriter.Write(b)
}

func (sw *SessionResponseWriter) IsWritten() bool {
	return sw.written
}

func (sw *SessionResponseWriter) WriteHeader(code int) {
	if !sw.written {
		if sw.CommitAndWriteSession != nil {
			sw.CommitAndWriteSession(sw.ResponseWriter, sw.Request)
		}
		sw.written = true
	}

	sw.ResponseWriter.WriteHeader(code)
}

func (sw *SessionResponseWriter) Unwrap() http.ResponseWriter {
	return sw.ResponseWriter
}
