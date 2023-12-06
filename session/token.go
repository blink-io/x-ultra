package session

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"github.com/teris-io/shortid"
)

func UUIDTokenGen() (string, error) {
	return uuid.NewString(), nil
}

func KSUIDTokenGen() (string, error) {
	return ksuid.New().String(), nil
}

func ShortIDTokenGen() (string, error) {
	return shortid.Generate()
}

func ShortUUIDTokenGen() (string, error) {
	return shortuuid.New(), nil
}

func XIDTokenGen() (string, error) {
	return xid.New().String(), nil
}
