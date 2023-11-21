package cbor

import (
	"bytes"

	"github.com/fxamacker/cbor/v2"
	"github.com/nats-io/nats.go"
)

const Name = "cbor"

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
	enc := cbor.NewEncoder(b)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (codec) Decode(subject string, data []byte, vPtr any) error {
	dec := cbor.NewDecoder(bytes.NewBuffer(data))
	err := dec.Decode(vPtr)
	return err
}

func (codec) Name() string {
	return Name
}
