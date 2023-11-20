package i18n

import (
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"time"
)

const (
	HTTPTimeout = 10 * time.Second
)

type (
	Bundler interface {
		LoadMessageFile(string) (*MessageFile, error)
		LoadMessageFileFS(fs.FS, string) (*MessageFile, error)
		LoadMessageFileBytes([]byte, string) (*MessageFile, error)
	}
	Loader interface {
		Load(Bundler) error
	}
)

var _ Bundler = (*Bundle)(nil)

// dirLoader loads files from directory
type dirLoader struct {
	dir string
}

func NewDirLoader(dir string) Loader {
	l := &dirLoader{dir: dir}
	return l
}

func (l *dirLoader) Load(b Bundler) error {
	if len(l.dir) == 0 {
		return nil
	}
	return filepath.WalkDir(l.dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// We cannot return fs.SkipDir because it won't walk into dir next
			return nil
		}
		// Ignore error?
		if _, err := b.LoadMessageFile(path); err != nil {
			log("[WARN] unable to load message from file: %s, reason: %+v", path, err)
		}
		return nil
	})
}

// fsLoader loads from embed.FS
type fsLoader struct {
	fs   fs.FS
	root string
}

func NewFSLoader(fs fs.FS, root string) Loader {
	l := &fsLoader{fs: fs, root: root}
	return l
}

func (l *fsLoader) Load(b Bundler) error {
	return fs.WalkDir(l.fs, l.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// Ignore error?
		if _, err := b.LoadMessageFileFS(l.fs, path); err != nil {
			log("[WARN] unable to load message from FS: %s", path)
		}
		return nil
	})
}

// httpLoader loads by http GET requests
// URL should be like: https://xxx.com/languages/zh_Hans.toml
type httpLoader struct {
	c       *http.Client
	extract func(string) string
	url     string
}

func NewHTTPLoader(url string, extract func(string) string, timeout time.Duration) Loader {
	if timeout == 0 {
		timeout = HTTPTimeout
	}
	c := &http.Client{Timeout: timeout}
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
		log("[WARN] unable to load message from URL: %s", h.url)
	}
	return nil
}

type bytesLoader struct {
	Path string
	Data []byte
}

func NewBytesLoader(path string, data []byte) Loader {
	return &bytesLoader{path, data}
}

func (h *bytesLoader) Load(b Bundler) error {
	if len(h.Path) > 0 && len(h.Data) > 0 {
		if _, err := b.LoadMessageFileBytes(h.Data, h.Path); err != nil {
			log("[WARN] unable to load message from bytes: %s", h.Path)
		}
	}
	return nil
}
