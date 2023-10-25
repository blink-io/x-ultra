package nestedtext

import (
	"bytes"

	"github.com/blink-io/x/encoding"

	"github.com/npillmayer/nestext"
	"github.com/npillmayer/nestext/ntenc"
)

const (
	Name = "nestedtext"
)

type codec struct {
}

var _ encoding.Codec = (*codec)(nil)

func New() encoding.Codec {
	return &codec{}
}

func (c *codec) Marshal(v any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := ntenc.Encode(v, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *codec) Unmarshal(data []byte, v any) error {
	r, err := nestext.Parse(bytes.NewReader(data))
	if err != nil {
		return err
	}
	//if rm, ok := r.(map[string]any); ok {
	//	for k, v := range rm {
	//		v[k] = m
	//	}
	//}
	v = r
	return nil
}

func (c *codec) Name() string {
	return Name
}
