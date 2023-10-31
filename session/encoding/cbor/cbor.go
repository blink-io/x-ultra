package cbor

import (
	"bytes"
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/fxamacker/cbor/v2"
)

const Name = "cbor"

type codec struct {
}

func New() encoding.Codec {
	return &codec{}
}

func (codec) Encode(deadline time.Time, values map[string]any) ([]byte, error) {
	pyd := &encoding.Payload{
		Deadline: deadline,
		Values:   values,
	}

	var b bytes.Buffer
	if err := cbor.NewEncoder(&b).Encode(pyd); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (codec) Decode(b []byte) (deadline time.Time, values map[string]any, err error) {
	pyd := new(encoding.Payload)

	r := bytes.NewReader(b)
	if err := cbor.NewDecoder(r).Decode(pyd); err != nil {
		return time.Time{}, nil, err
	}

	return pyd.Deadline, pyd.Values, nil
}

func (codec) Name() string {
	return Name
}
