package i18n

import (
	"io"
	"net/http"
	"time"
)

// httpLoader loads by http GET requests
// URLVar should be like: https://xxx.com/languages/zh_Hans.toml
type httpLoader struct {
	c       *http.Client
	extract func(string) string
	url     string
}

func NewHTTPLoader(url string, extract func(string) string, timeout time.Duration) Loader {
	if timeout == 0 {
		timeout = DefaultTimeout
	}
	c := &http.Client{Timeout: timeout}
	return NewHTTPClientLoader(c, url, extract)
}

func NewHTTPClientLoader(c *http.Client, url string, extract func(string) string) Loader {
	return &httpLoader{c: c, extract: extract, url: url}
}

func (h *httpLoader) Load(b Bundler) error {
	res, err := h.c.Get(h.url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var path string
	if h.extract != nil {
		path = h.extract(h.url)
	} else {
		path = h.url
	}
	if _, err := b.LoadMessageFileBytes(buf, path); err != nil {
		log("[WARN] unable to load message from URLVar: %s", h.url)
	}
	return nil
}
