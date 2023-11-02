package shared

import (
	"errors"
)

var (
	// ErrMultipleEndpointsUnsupported is thrown when there are
	// multiple endpoints specified for Redis.
	ErrMultipleEndpointsUnsupported = errors.New("redis: does not support multiple endpoints")

	// ErrAbortTryLock is thrown when a user stops trying to seek the lock
	// by sending a signal to the stop chan,
	// this is used to verify if the operation succeeded.
	ErrAbortTryLock = errors.New("redis: lock operation aborted")

	// ErrMasterSetMustBeProvided is thrown when Redis Sentinel is enabled
	// and the MasterName option is undefined.
	ErrMasterSetMustBeProvided = errors.New("master set name must be provided")

	// ErrInvalidRoutesOptions is thrown when Redis Sentinel is enabled
	// with RouteByLatency & RouteRandomly options without the ClusterClient.
	ErrInvalidRoutesOptions = errors.New("RouteByLatency and RouteRandomly options are only allowed with the ClusterClient")
)
