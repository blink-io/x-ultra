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
	// bundle is default bundle
	bundle = New(DefaultOptions)
	cc     = ttlcache.New[string, T](
		ttlcache.WithTTL[string, T](ttlcache.DefaultTTL),
	)
	fm  = sprig.TxtFuncMap()
	log = stdlog.Printf
)

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
	i, _ := cc.GetOrSet(lang, L(i18n.NewLocalizer(bundle.ib, lang)))
	return i.Value()
}

func Languages() []string {
	return bundle.Languages()
}

func LoadFromDir(root string) error {
	return NewDirLoader(root).Load(bundle)
}

func LoadFromFS(fs fs.FS, root string) error {
	return NewFSLoader(fs, root).Load(bundle)
}

func LoadFromHTTP(url string) error {
	return NewHTTPLoader(url, 10*time.Second).Load(bundle)
}

func PluralCount(pluralCount interface{}) LOption {
	return func(config *LocalizeConfig) {
		config.PluralCount = pluralCount
	}
}

// L defines Localizer wrapper function for translation
func L(loc *Localizer) T {
	return func(messageID string, ops ...LOption) string {
		if loc != nil {
			conf := &LocalizeConfig{
				MessageID: messageID,
				Funcs:     fm,
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
	return func(c *LocalizeConfig) {
		c.TemplateData = d
	}
}

func D(d map[string]any) LOption {
	return func(c *LocalizeConfig) {
		c.TemplateData = d
	}
}
