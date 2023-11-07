package msgpack

import (
	"go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
)

var _ converter.PayloadConverter = (*payloadConverter)(nil)

type payloadConverter struct {
}

func (p *payloadConverter) ToPayload(value interface{}) (*common.Payload, error) {
	//TODO implement me
	panic("implement me")
}

func (p *payloadConverter) FromPayload(payload *common.Payload, valuePtr interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (p *payloadConverter) ToString(payload *common.Payload) string {
	//TODO implement me
	panic("implement me")
}

func (p *payloadConverter) Encoding() string {
	//TODO implement me
	panic("implement me")
}
