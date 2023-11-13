package i18n

import (
	"io/fs"
	stdlog "log"
	"sync"
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
	MessageFile    = i18n.MessageFile
	Localizer      = i18n.Localizer
	LocalizeConfig = i18n.LocalizeConfig

	LOption func(*LocalizeConfig)
	// T is short for translation function
	T func(string, ...LOption) string

	MD map[string]any

	Bundle struct {
		*ib
	}

	Logger func(string, ...any)
)

var (
	globalMux sync.Mutex
	// bb is default bb
	bb = New(DefaultOptions)
	cc = ttlcache.New[string, T](
		ttlcache.WithTTL[string, T](ttlcache.DefaultTTL),
	)
	fm  = sprig.TxtFuncMap()
	log = stdlog.Printf
)

func New(o *Options) *Bundle {
	o = setupOptions(o)
	lang := o.Language
	if lang == language.Und {
		if l, err := locale.Detect(); err != nil {
			lang = language.English
		} else {
			lang = l
		}
	}

	ib := i18n.NewBundle(lang)
	for k, f := range unmarshalFns {
		ib.RegisterUnmarshalFunc(k, f)
	}

	b := &Bundle{ib}
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
	bb = b
	globalMux.Unlock()
}

func SetLogger(l Logger) {
	if log != nil {
		log = l
	}
}

func (b *Bundle) Clone() *Bundle {
	return b
}

func (b *Bundle) Languages() []string {
	var langs []string
	for _, t := range b.LanguageTags() {
		langs = append(langs, t.String())
	}
	return langs
}

func GetT(lang string) T {
	i, _ := cc.GetOrSet(lang, L(i18n.NewLocalizer(bb.ib, lang)))
	return i.Value()
}

func Languages() []string {
	return bb.Languages()
}

func LoadFromDir(root string) error {
	return NewDirLoader(root).Load(bb)
}

func LoadFromFS(fs fs.FS, root string) error {
	return NewFSLoader(fs, root).Load(bb)
}

func LoadFromHTTP(url string) error {
	return NewHTTPLoader(url, 10*time.Second).Load(bb)
}

func LoadFromBytes(path string, data []byte) error {
	return NewBytesLoader(path, data).Load(bb)
}
