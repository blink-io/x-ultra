// Package toml implements a koanf.Parser that parses TOML bytes as conf maps.
package toml

import (
	"github.com/pelletier/go-toml/v2"
)

// TOML implements a TOML parser.
type TOML struct{}

// Parser returns a TOML Parser.
func Parser() *TOML {
	return &TOML{}
}

// Unmarshal parses the given TOML bytes.
func (p *TOML) Unmarshal(b []byte) (map[string]interface{}, error) {
	d := make(map[string]interface{})
	err := toml.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Marshal marshals the given config map to TOML bytes.
func (p *TOML) Marshal(o map[string]interface{}) ([]byte, error) {
	return toml.Marshal(o)
}
