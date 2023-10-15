package i18n

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type (
	Loader interface {
		Load(*Bundle) error
	}
)

// dirLoader loads files from directory
type dirLoader struct {
	root     string
	suffixes []string
}

func NewDirLoader(root string, suffixes ...string) Loader {
	l := &dirLoader{root: root, suffixes: suffixes}
	return l
}

func (l *dirLoader) Load(b *Bundle) error {
	if len(l.root) == 0 {
		return nil
	}
	return filepath.WalkDir(l.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// We cannot return fs.SkipDir because it won't walk into dir next
			return nil
		}
		// Ignore unsupported suffixes
		if !supportedSuffixes(path, l.suffixes...) {
			log("unsupported suffix for path %s", path)
			return nil
		}
		// Ignore error?
		if _, err := b.LoadMessageFile(path); err != nil {
			log("unable to load message from file: %s", path)
		} else {
			log("loaded from file: %s", path)
		}
		return nil
	})
}

// fsLoader loads from embed.FS
type fsLoader struct {
	fs       fs.FS
	root     string
	suffixes []string
}

func NewFSLoader(fs fs.FS, root string, suffixes ...string) Loader {
	l := &fsLoader{fs: fs, root: root, suffixes: suffixes}
	return l
}

func (l *fsLoader) Load(b *Bundle) error {
	return fs.WalkDir(l.fs, l.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// Ignore unsupported suffixes
		if !supportedSuffixes(path, l.suffixes...) {
			log("unsupported suffix for path %s", path)
			return nil
		}
		// Ignore error?
		if _, err := b.LoadMessageFileFS(l.fs, path); err != nil {
			log("unable to load message from FS: %s", path)
		} else {
			log("loaded from file: %s", path)
		}
		return nil
	})
}

// httpLoader loads by http GET requests
type httpLoader struct {
	c   *http.Client
	url string
}

func NewHTTPLoader(url string, timeout time.Duration) Loader {
	c := &http.Client{Timeout: timeout}
	return &httpLoader{c, url}
}

func (h *httpLoader) Load(b *Bundle) error {
	//res, err := h.c.Get(h.url)
	//if err != nil {
	//	return err
	//}
	//
	//defer res.Body.Close()
	//
	//buf, err := io.ReadAll(res.Body)
	//if err != nil {
	//	return err
	//}
	//
	//b.ParseMessageFileBytes()
	return nil
}

func supportedSuffixes(path string, suffixes ...string) bool {
	for _, s := range suffixes {
		if strings.HasSuffix(path, s) {
			return true
		}
	}
	return false
}
