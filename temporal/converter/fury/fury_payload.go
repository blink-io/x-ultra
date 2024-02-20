package fury

import (
	"fmt"

	"github.com/apache/incubator-fury/go/fury"
	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
)

const (
	Name = "fury"

	MetadataEncoding = "binary/fury"
)

var _ converter.PayloadConverter = (*payloadConverter)(nil)

type payloadConverter struct {
}

func (c *payloadConverter) ToPayload(value interface{}) (*commonpb.Payload, error) {
	data, err := fury.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", converter.ErrUnableToEncode, err)
	}
	return newPayload(data, c), nil
}

func (c *payloadConverter) FromPayload(payload *commonpb.Payload, valuePtr interface{}) error {
	err := fury.Unmarshal(payload.GetData(), valuePtr)
	if err != nil {
		return fmt.Errorf("%w: %v", converter.ErrUnableToDecode, err)
	}
	return nil
}

func (c *payloadConverter) ToString(payload *commonpb.Payload) string {
	return string(payload.GetData())
}

func (c *payloadConverter) Encoding() string {
	return MetadataEncoding
}

func newPayload(data []byte, c converter.PayloadConverter) *commonpb.Payload {
	return &commonpb.Payload{
		Metadata: map[string][]byte{
			converter.MetadataEncoding: []byte(c.Encoding()),
		},
		Data: data,
	}
}

func (c *payloadConverter) Name() string {
	return Name
}
