package requestid

import (
	"github.com/google/uuid"
)

const (
	HeaderXRequestID = "X-Request-ID"
)

type Options struct {
	Header    string
	Generator func() string
}

var DefaultOptions = &Options{
	Header:    HeaderXRequestID,
	Generator: defaultGenerator,
}

func setupOptions(c *Options) *Options {
	if c == nil {
		return DefaultOptions
	}
	if c.Header == "" {
		c.Header = HeaderXRequestID
	}
	if c.Generator == nil {
		c.Generator = defaultGenerator
	}
	return c
}

func defaultGenerator() string {
	return uuid.New().String()
}
