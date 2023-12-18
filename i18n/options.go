package i18n

import (
	"golang.org/x/text/language"
)

type Options struct {
	Cache    bool
	Language language.Tag
	Loaders  []Loader
}

var (
	DefaultHeader  = "X-Language"
	DefaultDir     = "./locales"
	DefaultLoader  = NewDirLoader(DefaultDir)
	DefaultOptions = &Options{
		Cache: true,
		Loaders: []Loader{
			DefaultLoader,
		},
	}
)

func setupOptions(o *Options) *Options {
	if o == nil {
		o = DefaultOptions
	}
	return o
}
