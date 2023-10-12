package msgpack

import (
	"github.com/blink-io/x/encoding"

	"github.com/vmihailenco/msgpack/v5"
)

const (
	Name = "msgpack"
)

type codec struct {
}

func New() encoding.Codec {
	return &codec{}
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

func (c *codec) Name() string {
	return Name
}
