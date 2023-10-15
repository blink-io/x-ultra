package locales

import (
	"embed"
)

//go:embed *.toml
var EmbedFS embed.FS
