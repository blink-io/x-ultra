package render

import (
	"errors"
	"io"
	"net/http"

	"github.com/unrolled/render"
	"google.golang.org/protobuf/proto"
)

const (
	// ContentProtobuf header value for Text data.
	ContentProtobuf = "application/x-protobuf"

	NameProtobuf = "protobuf"
)

type Protobuf struct {
	render.Head
}

// Render a Protobuf response.
func (t Protobuf) Render(w io.Writer, v interface{}) error {
	if hw, ok := w.(http.ResponseWriter); ok {
		c := hw.Header().Get(render.ContentType)
		if c != "" {
			t.Head.ContentType = c
		}
		t.Head.Write(hw)
	}

	m, ok := v.(proto.Message)
	if !ok {
		return errors.New("v is not a protobuf message")
	}

	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	_, _ = w.Write(data)
	return nil
}

func (Protobuf) Name() string {
	return NameProtobuf
}
