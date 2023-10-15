package render

import (
	"io"
	"net/http"

	"github.com/unrolled/render"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	// ContentMsgpack header value for Msgpack data.
	ContentMsgpack = "application/msgpack; charset=utf-8"

	NameMsgpack = "msgpack"
)

type Msgpack struct {
	render.Head
}

// Render a Msgpack response.
func (t Msgpack) Render(w io.Writer, v interface{}) error {
	if hw, ok := w.(http.ResponseWriter); ok {
		c := hw.Header().Get(render.ContentType)
		if c != "" {
			t.Head.ContentType = c
		}
		t.Head.Write(hw)
	}

	data, err := msgpack.Marshal(v)
	if err != nil {
		return err
	}

	_, _ = w.Write(data)
	return nil
}

func (Msgpack) Name() string {
	return NameMsgpack
}
