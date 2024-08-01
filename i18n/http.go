package i18n

import (
	"io"
	"net/http"
)

// httpLoader loads by http GET requests
// URLVar should be like: https://xxx.com/languages/zh_Hans.toml
type httpLoader struct {
	url  string
	opts *httpOptions
}

func NewHTTPLoader(url string, ops ...HTTPOption) Loader {
	opts := applyHTTPOptions(ops...)
	return &httpLoader{url: url, opts: opts}
}

func (h *httpLoader) Load(b Bundler) error {
	c := h.opts.httpClient
	rf := h.opts.requestFunc

	var res *http.Response
	var err error
	if rf != nil {
		if req, rerr := h.opts.requestFunc(c, h.url); rerr != nil {
			return rerr
		} else {
			res, err = c.Do(req)
		}
	} else {
		res, err = c.Get(h.url)
	}
	if err != nil {
		return err
	}

	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var path string
	if ef := h.opts.extractFunc; ef != nil {
		path = ef(h.url)
	} else {
		path = h.url
	}
	if _, err := b.LoadMessageFileBytes(buf, path); err != nil {
		log("[WARN] unable to load message from URL: %s", h.url)
	}
	return nil
}

type httpOptions struct {
	httpClient *http.Client

	requestFunc func(*http.Client, string) (*http.Request, error)

	extractFunc func(string) string
}

type HTTPOption func(*httpOptions)

func applyHTTPOptions(ops ...HTTPOption) *httpOptions {
	opts := &httpOptions{
		httpClient: &http.Client{Timeout: DefaultTimeout},
		extractFunc: func(s string) string {
			return s
		},
	}
	for _, op := range ops {
		op(opts)
	}
	return opts
}

func WithHTTPClient(httpClient *http.Client) HTTPOption {
	return func(o *httpOptions) {
		o.httpClient = httpClient
	}
}

func WithRequestFunc(requestFunc func(*http.Client, string) (*http.Request, error)) HTTPOption {
	return func(o *httpOptions) {
		o.requestFunc = requestFunc
	}
}

func WithExtractFunc(extractFunc func(string) string) HTTPOption {
	return func(o *httpOptions) {
		o.extractFunc = extractFunc
	}
}
