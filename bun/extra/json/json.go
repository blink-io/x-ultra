package json

import (
	"io"

	"github.com/goccy/go-json"
	"github.com/uptrace/bun/extra/bunjson"
)

var _ bunjson.Provider = (*provider)(nil)

type provider struct{}

func NewProvider() bunjson.Provider {
	return &provider{}
}

func (provider) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (provider) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (provider) NewEncoder(w io.Writer) bunjson.Encoder {
	return json.NewEncoder(w)
}

func (provider) NewDecoder(r io.Reader) bunjson.Decoder {
	return json.NewDecoder(r)
}
