package render

import (
	"io"
	"sync"

	"github.com/unrolled/render"
)

type Marshaler func(v interface{}) ([]byte, error)

type rr = render.Render

type Render struct {
	*rr
}

type Options = render.Options

var once sync.Once

var def *Render

func Default() *Render {
	once.Do(func() {
		def = New()
	})
	return def
}
func New(o ...Options) *Render {
	rr := render.New(o...)
	return &Render{rr}
}

// Protobuf marshals the given interface object and writes the Protobuf response.
func (r *Render) Protobuf(w io.Writer, status int, v interface{}) error {
	head := render.Head{
		ContentType: ContentProtobuf,
		Status:      status,
	}

	j := Protobuf{
		Head: head,
	}

	return r.rr.Render(w, j, v)
}

// TOML marshals the given interface object and writes the TOML response.
func (r *Render) TOML(w io.Writer, status int, v interface{}) error {
	head := render.Head{
		ContentType: ContentTOML,
		Status:      status,
	}

	e := TOML{
		Head: head,
	}

	return r.rr.Render(w, e, v)
}

// YAML marshals the given interface object and writes the YAML response.
func (r *Render) YAML(w io.Writer, status int, v interface{}) error {
	head := render.Head{
		ContentType: ContentYAML,
		Status:      status,
	}

	e := YAML{
		Head: head,
	}

	return r.rr.Render(w, e, v)
}

// Msgpack marshals the given interface object and writes the Msgpack response.
func (r *Render) Msgpack(w io.Writer, status int, v interface{}) error {
	head := render.Head{
		ContentType: ContentMsgpack,
		Status:      status,
	}

	e := Msgpack{
		Head: head,
	}

	return r.rr.Render(w, e, v)
}

// Cbor marshals the given interface object and writes the Cbor response.
func (r *Render) Cbor(w io.Writer, status int, v interface{}) error {
	head := render.Head{
		ContentType: ContentCbor,
		Status:      status,
	}

	e := Cbor{
		Head: head,
	}

	return r.rr.Render(w, e, v)
}
