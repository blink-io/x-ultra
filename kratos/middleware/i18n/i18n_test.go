package i18n

import (
	"github.com/blink-io/x/i18n"

	"github.com/pelletier/go-toml/v2"
)

var bundle = i18n.Default()

func init() {
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
}
