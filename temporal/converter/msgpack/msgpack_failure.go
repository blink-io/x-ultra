package msgpack

import (
	"go.temporal.io/api/failure/v1"
	"go.temporal.io/sdk/converter"
)

var _ converter.FailureConverter = (*failureConverter)(nil)

type failureConverter struct {
}

func (f *failureConverter) ErrorToFailure(err error) *failure.Failure {
	//TODO implement me
	panic("implement me")
}

func (f *failureConverter) FailureToError(failure *failure.Failure) error {
	//TODO implement me
	panic("implement me")
}
