package i18n

type (
	Options struct {
		Loaders []Loader
	}
)

var (
	DefaultHeader  = "X-Language"
	DefaultDir     = "./locales"
	DefaultLoader  = NewDirLoader(DefaultDir)
	DefaultOptions = &Options{
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
