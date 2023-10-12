package tsz

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

var _ bun.BeforeAppendModelHook = (*Model)(nil)

type Model struct {
	CreatedAt time.Time    `bun:"created_at,nullzero,notnull" db:"created_at" json:"created_at,omitempty" toml:"created_at,omitempty" yaml:"created_at,omitempty" msgpack:"created_at,omitempty"`
	UpdatedAt time.Time    `bun:"updated_at,nullzero,notnull" db:"updated_at" json:"updated_at,omitempty" toml:"updated_at,omitempty" yaml:"updated_at,omitempty" msgpack:"updated_at,omitempty"`
	DeletedAt bun.NullTime `bun:"deleted_at" db:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at,omitempty" yaml:"deleted_at,omitempty" msgpack:"deleted_at,omitempty"`
}

func (m *Model) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.Operation() {
	case "INSERT":
		m.CreatedAt = time.Now()
		m.UpdatedAt = m.CreatedAt
	case "UPDATE":
		m.UpdatedAt = time.Now()
	case "DELETE":
		m.DeletedAt = bun.NullTime{Time: time.Now()}
	}
	return nil
}
