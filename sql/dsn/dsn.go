package dsn

import (
	"net/url"
)

// Parser parses DSN string to url.URL
type Parser interface {
	Parse(string) (*url.URL, error)
}

// Converter converts url.URL to DNS string
type Converter interface {
	Convert(*url.URL) (string, error)
}

type Processor interface {
	Parser
	Converter
	Name() string
}
