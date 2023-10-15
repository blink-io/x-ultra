package render

import (
	"io"
	"net/http"

	"github.com/unrolled/render"
	"gopkg.in/yaml.v3"
)

const (
	// ContentYAML header value for YAML format data.
	ContentYAML = "application/yaml"

	NameYAML = "yaml"
)

type YAML struct {
	render.Head
}

// Render a YAML format response.
func (t YAML) Render(w io.Writer, v interface{}) error {
	if hw, ok := w.(http.ResponseWriter); ok {
		c := hw.Header().Get(render.ContentType)
		if c != "" {
			t.Head.ContentType = c
		}
		t.Head.Write(hw)
	}

	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	_, _ = w.Write(data)
	return nil
}

func (YAML) Name() string {
	return NameYAML
}
