package fury

import (
	"time"

	"github.com/blink-io/x/session/encoding"

	"github.com/apache/incubator-fury/go/fury"
)

const Name = "fury"

func init() {
	encoding.Register(Name, &codec{})
}

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

	return fury.Marshal(pyd)
}

func (codec) Decode(data []byte) (deadline time.Time, values map[string]any, err error) {
	pyd := new(encoding.Payload)

	if err := fury.Unmarshal(data, pyd); err != nil {
		return time.Time{}, nil, err
	}

	return pyd.Deadline, pyd.Values, nil
}

func (codec) Name() string {
	return Name
}
