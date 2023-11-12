package json

import (
	"bytes"

	"github.com/nats-io/nats.go"
	"github.com/vmihailenco/msgpack/v5"
)

const Name = "msgpack"

func init() {
	nats.RegisterEncoder(Name, New())
}

type codec struct {
}

func New() nats.Encoder {
	return &codec{}
}

func (codec) Encode(subject string, v any) ([]byte, error) {
	b := new(bytes.Buffer)
	enc := msgpack.NewEncoder(b)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (codec) Decode(subject string, data []byte, vPtr any) error {
	dec := msgpack.NewDecoder(bytes.NewBuffer(data))
	err := dec.Decode(vPtr)
	return err
}

func (codec) Name() string {
	return Name
}
