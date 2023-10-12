package bun

import (
	"database/sql"

	"github.com/blink-io/x/bun/model/mixin"
	"github.com/uptrace/bun"
)

type mxm = mixin.Model

// Application represents iOS/Android/Windows/OSX/Linux application
type Application struct {
	bun.BaseModel `bun:"applications,alias:applications" db:"-" json:"-" toml:"-" yaml:"-" msgpack:"-"`
	mxm
	Status      string         `bun:"status,notnull" db:"status" json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" msgpack:"status,omitempty"`
	Type        string         `bun:"type,notnull" db:"type" json:"type,omitempty" toml:"type,omitempty" yaml:"type,omitempty" msgpack:"type,omitempty"`
	Name        string         `bun:"name,notnull" db:"name" json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" msgpack:"name,omitempty"`
	Code        string         `bun:"code,unique,notnull" db:"code" json:"code,omitempty" toml:"code,omitempty" yaml:"code,omitempty" msgpack:"code,omitempty"`
	Description sql.NullString `bun:"description" db:"description" json:"description,omitempty" toml:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
}
