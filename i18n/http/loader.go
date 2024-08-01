package http

import (
	"github.com/blink-io/x/i18n"
	"io"
	"net/http"
)

// httpLoader loads by http GET requests
// URLVar should be like: https://xxx.com/languages/zh_Hans.toml
type httpLoader struct {
	url  string
	opts *options
}

func NewHTTPLoader(url string, ops ...Option) i18n.Loader {
	opts := applyHTTPOptions(ops...)
	return &httpLoader{url: url, opts: opts}
}

func LoadFromHTTP(url string, ops ...Option) error {
	return NewHTTPLoader(url, ops...).Load(i18n.Default())
}

func (h *httpLoader) Load(b i18n.Bundler) error {
	c := h.opts.client

	var res *http.Response
	var err error
	if rf := h.opts.requestFunc; rf != nil {
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
		i18n.GetLogger()("[WARN] unable to load message from URL: %s", h.url)
	}
	return nil
}
