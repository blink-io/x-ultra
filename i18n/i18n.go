package i18n

import (
	"fmt"
	"io/fs"
	"log/slog"
	"sync"
	"text/template"
	"time"

	"github.com/blink-io/x/locale"

	"github.com/go-task/slim-sprig/v3"
	"github.com/jellydator/ttlcache/v3"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type (
	// ib is short for i18n.Bundle
	ib             = i18n.Bundle
	Message        = i18n.Message
	MessageFile    = i18n.MessageFile
	Localizer      = i18n.Localizer
	LocalizeConfig = i18n.LocalizeConfig
	UnmarshalFunc  = i18n.UnmarshalFunc

	LOption func(*LocalizeConfig)
	// T is short for translation function
	T func(messageID string, ops ...LOption) string

	MD map[string]any

	Bundle struct {
		*ib
		cache   *ttlcache.Cache[string, T]
		noCache bool
		funcMap template.FuncMap
	}
)

const (
	defaultTTL = 2 * time.Hour
)

var (
	globalMux sync.Mutex
	// bb is default bb
	bb = New(DefaultOptions)

	fm         = sprig.TxtFuncMap()
	log Logger = func(format string, args ...any) {
		msg := fmt.Sprintf(format, args...)
		slog.Info(msg)
	}
)

func New(o *Options) *Bundle {
	o = setupOptions(o)
	lang := o.Language
	if lang == language.Und {
		if l, err := locale.Detect(); err != nil {
			// Use English as default
			lang = language.English
		} else {
			lang = l
		}
	}

	ib := i18n.NewBundle(lang)
	for k, f := range unmarshalFns {
		ib.RegisterUnmarshalFunc(k, f)
	}

	b := &Bundle{
		ib:      ib,
		noCache: o.Cache,
		funcMap: fm,
	}
	if !b.noCache {
		b.cache = ttlcache.New[string, T](
			ttlcache.WithTTL[string, T](defaultTTL),
		)
	}
	for _, l := range o.Loaders {
		_ = l.Load(b)
	}

	return b
}

// Default gets default Bundle
func Default() *Bundle {
	return bb
}

// Replace replaces default Bundle
func Replace(b *Bundle) {
	globalMux.Lock()
	defer globalMux.Unlock()
	bb = b
}

func (b *Bundle) LoadMessageFileBytes(buf []byte, path string) (*MessageFile, error) {
	return b.ParseMessageFileBytes(buf, path)
}

func (b *Bundle) Languages() []string {
	var langs []string
	for _, t := range b.LanguageTags() {
		langs = append(langs, t.String())
	}
	return langs
}

func (b *Bundle) Load(l Loader) error {
	return l.Load(b)
}

func (b *Bundle) LoadFromDir(dir string) error {
	return NewDirLoader(dir).Load(b)
}

func (b *Bundle) LoadFromFS(fs fs.FS, root string) error {
	return NewFSLoader(fs, root).Load(b)
}

func (b *Bundle) LoadFromHTTP(url string, extract func(string) string) error {
	return NewHTTPLoader(url, extract, DefaultTimeout).Load(b)
}

func (b *Bundle) LoadFromBytes(path string, data []byte) error {
	return NewBytesLoader(path, data).Load(b)
}

func GetT(lang string) T {
	return bb.T(lang)
}

func (b *Bundle) T(lang string) T {
	return doT(b, lang)
}

func doT(b *Bundle, lang string) T {
	if b.noCache {
		tr := Tr(i18n.NewLocalizer(b.ib, lang))
		return tr
	}
	if i := b.cache.Get(lang); i != nil && !i.IsExpired() {
		return i.Value()
	} else {
		tr := Tr(i18n.NewLocalizer(b.ib, lang))
		b.cache.Set(lang, tr, defaultTTL)
		return tr
	}
}

func Languages() []string {
	return bb.Languages()
}

func LoadFromDir(dir string) error {
	return NewDirLoader(dir).Load(bb)
}

func LoadFromFS(fs fs.FS, root string) error {
	return NewFSLoader(fs, root).Load(bb)
}

func LoadFromHTTP(url string, extract func(string) string) error {
	return NewHTTPLoader(url, extract, DefaultTimeout).Load(bb)
}

func LoadFromBytes(path string, data []byte) error {
	return NewBytesLoader(path, data).Load(bb)
}
