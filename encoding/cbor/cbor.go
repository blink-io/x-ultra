package cbor

import (
	"github.com/blink-io/x/encoding"

	"github.com/fxamacker/cbor/v2"
)

const (
	Name = "cbor"
)

type codec struct {
}

func New() encoding.Codec {
	return &codec{}
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	return cbor.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	return cbor.Unmarshal(data, v)
}

func (c *codec) Name() string {
	return Name
}
