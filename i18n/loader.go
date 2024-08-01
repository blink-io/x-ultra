package i18n

import (
	"io/fs"
	"path/filepath"
	"time"
)

const (
	DefaultTimeout = 10 * time.Second
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
