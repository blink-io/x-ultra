package render

import (
	"io"
	"net/http"

	"github.com/fxamacker/cbor/v2"
	"github.com/unrolled/render"
)

const (
	// ContentCbor header value for cbor data.
	ContentCbor = "application/cbor; charset=utf-8"

	NameCbor = "cbor"
)

type Cbor struct {
	render.Head
}

// Render a Msgpack response.
func (t Cbor) Render(w io.Writer, v interface{}) error {
	if hw, ok := w.(http.ResponseWriter); ok {
		c := hw.Header().Get(render.ContentType)
		if c != "" {
			t.Head.ContentType = c
		}
		t.Head.Write(hw)
	}

	data, err := cbor.Marshal(v)
	if err != nil {
		return err
	}

	_, _ = w.Write(data)
	return nil
}

func (Cbor) Name() string {
	return NameCbor
}
