// Package tomlv2 implements a koanf.Parser that parses TOML bytes as conf maps.
package tomlv2

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
func (p *TOML) Unmarshal(b []byte) (map[string]any, error) {
	d := make(map[string]any)
	err := toml.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Marshal marshals the given config map to TOML bytes.
func (p *TOML) Marshal(o map[string]any) ([]byte, error) {
	return toml.Marshal(o)
}
