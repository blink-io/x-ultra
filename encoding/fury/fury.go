package fury

import (
	"github.com/blink-io/x/encoding"

	"github.com/apache/incubator-fury/go/fury"
)

const (
	Name = "fury"
)

type codec struct {
}

func New() encoding.Codec {
	return &codec{}
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	return fury.Marshal(v)
}

func (c *codec) Unmarshal(data []byte, v interface{}) error {
	return fury.Unmarshal(data, v)
}

func (c *codec) Name() string {
	return Name
}
