package i18n

import (
	"github.com/goccy/go-json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

var unmarshalFns = map[string]i18n.UnmarshalFunc{
	"yaml": yaml.Unmarshal,
	"yml":  yaml.Unmarshal,
	"toml": toml.Unmarshal,
	"json": json.Unmarshal,
}

var DefaultSuffixes = []string{
	".yaml",
	".yml",
	".toml",
	".json",
}
