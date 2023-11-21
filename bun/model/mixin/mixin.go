package mixin

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var (
	// m is a lock for SetGenerator
	m = new(sync.Mutex)

	// gg is global generator for all tables' ID field
	gg = func() string {
		return uuid.New().String()
	}

	Columns = &columns{
		// Required fields
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		// Optional fields
		CreatedBy: "created_by",
		UpdatedBy: "updated_by",
		DeletedAt: "deleted_at",
		DeletedBy: "deleted_by",
		IsDeleted: "is_deleted",
	}
)

type columns struct {
	// Require columns
	ID        string
	CreatedAt string
	UpdatedAt string
	// Optional columns
	CreatedBy string
	UpdatedBy string
	DeletedAt string
	DeletedBy string
	IsDeleted string
}

type Generator func() string

// Model is the common part for all models in the project
var _ bun.BeforeAppendModelHook = (*Model)(nil)

type Model struct {
	ig        Generator // ID generator for a single model
	ID        string    `bun:"id,pk" db:"id" json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty" msgpack:"id,omitempty"`
	CreatedAt time.Time `bun:"created_at,notnull,skipupdate" db:"created_at" json:"created_at,omitempty" toml:"created_at,omitempty" yaml:"created_at,omitempty" msgpack:"created_at,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,notnull" db:"updated_at" json:"updated_at,omitempty" toml:"updated_at,omitempty" yaml:"updated_at,omitempty" msgpack:"updated_at,omitempty"`
	// Optional fields for tables
	CreatedBy sql.NullString `bun:"created_by,nullzero,skipupdate" db:"created_by" json:"created_by,omitempty" toml:"created_by,omitempty" yaml:"created_by,omitempty" msgpack:"created_by,omitempty"`
	UpdatedBy sql.NullString `bun:"updated_by,nullzero" db:"updated_by" json:"updated_by,omitempty" toml:"updated_by,omitempty" yaml:"updated_by,omitempty" msgpack:"updated_by,omitempty"`
	DeletedAt bun.NullTime   `bun:"deleted_at,nullzero,skipupdate" db:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at,omitempty" yaml:"deleted_at,omitempty" msgpack:"deleted_at,omitempty"`
	DeletedBy sql.NullString `bun:"deleted_by,nullzero,skipupdate" db:"deleted_by" json:"deleted_by,omitempty" toml:"deleted_by,omitempty" yaml:"deleted_by,omitempty" msgpack:"deleted_by,omitempty"`
	IsDeleted sql.NullBool   `bun:"is_deleted,nullzero,skipupdate" db:"is_deleted" json:"is_deleted,omitempty" toml:"is_deleted,omitempty" yaml:"is_deleted,omitempty" msgpack:"is_deleted,omitempty"`
}

func (m *Model) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	handleTSZ(m, query)
	handleAutoID(m, query)
	return nil
}

// SetGenerator overrides the global generator for every model
func (m *Model) SetGenerator(g Generator) {
	m.ig = g
}

// SetGenerator sets global ID generator for all models
func SetGenerator(g Generator) {
	if g != nil {
		m.Lock()
		gg = g
		m.Unlock()
	} else {
		log.Println("parameter g is nil, ignore")
	}
}

func handleTSZ(m *Model, query bun.Query) {
	if m != nil {
		switch query.Operation() {
		case "INSERT":
			m.CreatedAt = time.Now()
			m.UpdatedAt = m.CreatedAt
		case "UPDATE":
			m.UpdatedAt = time.Now()
		case "DELETE":
			m.DeletedAt = bun.NullTime{Time: time.Now()}
			m.IsDeleted = sql.NullBool{Bool: true, Valid: true}

		}
	}
}

func handleAutoID(m *Model, query bun.Query) {
	if o := query.Operation(); o == "INSERT" && m != nil && len(m.ID) == 0 {
		var xg = gg
		if m.ig != nil {
			xg = m.ig
		}
		m.ID = xg()
	}
}
