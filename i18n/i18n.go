package i18n

import (
	"io/fs"
	stdlog "log"
	"sync"
	"time"

	"github.com/Xuanwo/go-locale"
	"github.com/go-task/slim-sprig/v3"
	"github.com/jellydator/ttlcache/v3"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type (
	// ib is short for i18n.Bundle
	ib = i18n.Bundle

	LOption func(*i18n.LocalizeConfig)
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
	// bundle is default bundle
	bundle          = New(DefaultOptions)
	tc              = ttlcache.New[string, T]()
	log             = stdlog.Printf
	DefaultSuffixes = make([]string, 0)
)

func init() {
	for k := range unmarshalFns {
		DefaultSuffixes = append(DefaultSuffixes, k)
	}
}

func New(o *Options) *Bundle {
	o = setupOptions(o)
	lang, err := locale.Detect()
	if err != nil {
		lang = language.English
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

// Default gets default bundle
func Default() *Bundle {
	return bundle
}

// Replace replaces default bundle
func Replace(b *Bundle) {
	globalMux.Lock()
	bundle = b
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
	if t := tc.Get(lang); t != nil {
		return t.Value()
	} else {
		nt := L(i18n.NewLocalizer(bundle.ib, lang))
		tc.Set(lang, nt, ttlcache.DefaultTTL)
		return nt
	}
}

func Languages() []string {
	return bundle.Languages()
}

func LoadFromDir(root string) error {
	return NewDirLoader(root, DefaultSuffixes...).Load(bundle)
}

func LoadFromFS(fs fs.FS, root string) error {
	return NewFSLoader(fs, root, DefaultSuffixes...).Load(bundle)
}

func LoadFromHTTP(url string) error {
	return NewHTTPLoader(url, 10*time.Second).Load(bundle)
}

func PluralCount(pluralCount interface{}) LOption {
	return func(config *i18n.LocalizeConfig) {
		config.PluralCount = pluralCount
	}
}

// L defines Localizer wrapper function for translation
func L(loc *i18n.Localizer) T {
	return func(messageID string, ops ...LOption) string {
		if loc != nil {
			conf := &i18n.LocalizeConfig{
				MessageID: messageID,
				Funcs:     sprig.TxtFuncMap(),
			}
			for _, o := range ops {
				o(conf)
			}
			if s, err := loc.Localize(conf); err == nil {
				return s
			}
		}
		return messageID
	}
}

func (d MD) O() LOption {
	return func(c *i18n.LocalizeConfig) {
		c.TemplateData = d
	}
}

func D(d map[string]any) LOption {
	return func(c *i18n.LocalizeConfig) {
		c.TemplateData = d
	}
}
