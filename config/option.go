package config

// Option configures how config loads the configuration.
type Option func(f *config)

// File returns an option that configures the filename that config
// looks for to provide the config values.
//
// The name must include the extension of the file. Supported
// file types are `yaml`, `yml`, `json` and `toml`.
//
//	config.Load(&cfg, config.File("config.toml"))
//
// If this option is not used then config looks for a file with name `config.yaml`.
func File(name string) Option {
	return func(f *config) {
		f.filename = name
	}
}

// IgnoreFile returns an option which disables any file lookup.
//
// This option effectively renders any `File` and `Dir` options useless. This option
// is most useful in conjunction with the `UseEnv` option when you want to provide
// config values only via environment variables.
//
//	config.Load(&cfg, config.IgnoreFile(), config.UseEnv("my_app"))
func IgnoreFile() Option {
	return func(f *config) {
		f.ignoreFile = true
	}
}

// Dirs returns an option that configures the directories that config searches
// to find the configuration file.
//
// Directories are searched sequentially and the first one with a matching config file is used.
//
// This is useful when you don't know where exactly your configuration will be during run-time:
//
//	config.Load(&cfg, config.Dirs(".", "/etc/myapp", "/home/user/myapp"))
//
// If this option is not used then config looks in the directory it is run from.
func Dirs(dirs ...string) Option {
	return func(f *config) {
		f.dirs = dirs
	}
}

// Tag returns an option that configures the tag key that config uses
// when for the alt name struct tag key in fields.
//
//	config.Load(&cfg, config.Tag("config"))
//
// If this option is not used then config uses the tag `config`.
func Tag(tag string) Option {
	return func(f *config) {
		f.tag = tag
	}
}

// TimeLayout returns an option that configures the time layout that config uses when
// parsing a time in a config file or in the default tag for time.Time fields.
//
//	config.Load(&cfg, config.TimeLayout("2006-01-02"))
//
// If this option is not used then config parses times using `time.RFC3339` layout.
func TimeLayout(layout string) Option {
	return func(f *config) {
		f.timeLayout = layout
	}
}

// UseEnv returns an option that configures config to additionally load values
// from the environment.
//
//	config.Load(&cfg, config.UseEnv("my_app"))
//
// Values loaded from the environment overwrite values loaded by the config file (if any).
//
// config looks for environment variables in the format PREFIX_FIELD_PATH or
// FIELD_PATH if prefix is empty. Prefix is capitalized regardless of what
// is provided. The field's path is formed by prepending its name with the
// names of all surrounding fields up to the root struct. If a field has
// an alternative name defined inside a struct tag then that name is
// preferred.
//
//	type Config struct {
//	  Build    time.Time
//	  LogLevel string `config:"log_level"`
//	  Server   struct {
//	    Host string
//	  }
//	}
//
// With the struct above and UseEnv("myapp") config would search for the following
// environment variables:
//
//	MYAPP_BUILD
//	MYAPP_LOG_LEVEL
//	MYAPP_SERVER_HOST
func UseEnv(prefix string) Option {
	return func(f *config) {
		f.useEnv = true
		f.envPrefix = prefix
	}
}

// UseStrict returns an option that configures config to return an error if
// there exists additional fields in the config file that are not defined
// in the config struct.
//
//	config.Load(&cfg, config.UseStrict())
//
// If this option is not used then config ignores any additional fields in the config file.
func UseStrict() Option {
	return func(f *config) {
		f.useStrict = true
	}
}
