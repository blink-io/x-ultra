package msgpack

import (
	"go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
)

var _ converter.DataConverter = (*dataConverter)(nil)

type dataConverter struct {
}

func (d dataConverter) ToPayload(value interface{}) (*common.Payload, error) {
	//TODO implement me
	panic("implement me")
}

func (d dataConverter) FromPayload(payload *common.Payload, valuePtr interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d dataConverter) ToPayloads(value ...interface{}) (*common.Payloads, error) {
	//TODO implement me
	panic("implement me")
}

func (d dataConverter) FromPayloads(payloads *common.Payloads, valuePtrs ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d dataConverter) ToString(input *common.Payload) string {
	//TODO implement me
	panic("implement me")
}

func (d dataConverter) ToStrings(input *common.Payloads) []string {
	//TODO implement me
	panic("implement me")
}
