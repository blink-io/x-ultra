package i18n

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

var unmarshalFns = make(map[string]i18n.UnmarshalFunc)

var DefaultSuffixes = make([]string, 0)

func init() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	unmarshalFns["yaml"] = yaml.Unmarshal
	unmarshalFns["yml"] = yaml.Unmarshal
	unmarshalFns["json"] = json.Unmarshal
	unmarshalFns["toml"] = toml.Unmarshal

	for k := range unmarshalFns {
		DefaultSuffixes = append(DefaultSuffixes, k)
	}
}
