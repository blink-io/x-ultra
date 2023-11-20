package i18n

import (
	"github.com/goccy/go-json"
	"github.com/pelletier/go-toml/v2"
	"github.com/vmihailenco/msgpack/v5"
	"gopkg.in/yaml.v3"
)

var unmarshalFns = map[string]UnmarshalFunc{
	"yaml":    yaml.Unmarshal,
	"yml":     yaml.Unmarshal,
	"toml":    toml.Unmarshal,
	"json":    json.Unmarshal,
	"msgpack": msgpack.Unmarshal,
}
