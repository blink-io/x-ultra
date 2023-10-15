package render

import (
	"io"
	"net/http"

	"github.com/pelletier/go-toml/v2"
	"github.com/unrolled/render"
)

const (
	// ContentTOML header value for TOML format data.
	ContentTOML = "application/toml"

	NameTOML = "toml"
)

type TOML struct {
	render.Head
}

// Render a TOML format response.
func (t TOML) Render(w io.Writer, v interface{}) error {
	if hw, ok := w.(http.ResponseWriter); ok {
		c := hw.Header().Get(render.ContentType)
		if c != "" {
			t.Head.ContentType = c
		}
		t.Head.Write(hw)
	}

	data, err := toml.Marshal(v)
	if err != nil {
		return err
	}

	_, _ = w.Write(data)
	return nil
}

func (TOML) Name() string {
	return NameTOML
}
