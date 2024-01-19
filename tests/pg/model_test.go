package pg

import (
	"database/sql"

	"github.com/blink-io/x/bun/model"
	"github.com/uptrace/bun"
)

// Application represents iOS/Android/Windows/OSX/Linux application
type Application struct {
	bun.BaseModel `bun:"applications,alias:applications" db:"-" json:"-" toml:"-" yaml:"-" msgpack:"-"`
	model.IDModel
	Status      string         `bun:"status,type:varchar(60),notnull" db:"status" json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" msgpack:"status,omitempty"`
	Type        string         `bun:"type,type:varchar(60),notnull" db:"type" json:"type,omitempty" toml:"type,omitempty" yaml:"type,omitempty" msgpack:"type,omitempty"`
	Name        string         `bun:"name,type:varchar(200),notnull" db:"name" json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" msgpack:"name,omitempty"`
	Code        string         `bun:"code,type:varchar(60),unique,notnull" db:"code" json:"code,omitempty" toml:"code,omitempty" yaml:"code,omitempty" msgpack:"code,omitempty"`
	Description sql.NullString `bun:"description,type:text" db:"description" json:"description,omitempty" toml:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	model.ExtraModel
}

func (Application) TableName() string {
	return "applications"
}

func (Application) Table() string {
	return "applications"
}
