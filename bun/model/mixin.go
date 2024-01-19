package model

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
	// guidMux is a lock for SetGenerator
	guidMux = new(sync.Mutex)

	// gg is global generator for all tables' GUID field
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

// MixinModel is the common part for all models in the project
var _ bun.BeforeAppendModelHook = (*MixinModel)(nil)

type MixinModel struct {
	// ID generator for a single model
	gg        Generator `bun:"-" db:"-" json:"-" toml:"-" yaml:"-" msgpack:"-"`
	ID        int64     `bun:"id,pk,autoincrement" db:"id" json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty" msgpack:"id,omitempty"`
	GUID      string    `bun:"guid,type:varchar(60)" db:"guid" json:"guid,omitempty" toml:"guid,omitempty" yaml:"guid,omitempty" msgpack:"guid,omitempty"`
	CreatedAt time.Time `bun:"created_at,notnull,skipupdate" db:"created_at" json:"created_at,omitempty" toml:"created_at,omitempty" yaml:"created_at,omitempty" msgpack:"created_at,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,notnull" db:"updated_at" json:"updated_at,omitempty" toml:"updated_at,omitempty" yaml:"updated_at,omitempty" msgpack:"updated_at,omitempty"`
	// Optional fields for tables
	CreatedBy sql.NullString `bun:"created_by,type:varchar(60),nullzero,skipupdate" db:"created_by" json:"created_by,omitempty" toml:"created_by,omitempty" yaml:"created_by,omitempty" msgpack:"created_by,omitempty"`
	UpdatedBy sql.NullString `bun:"updated_by,type:varchar(60),nullzero" db:"updated_by" json:"updated_by,omitempty" toml:"updated_by,omitempty" yaml:"updated_by,omitempty" msgpack:"updated_by,omitempty"`
	DeletedAt sql.NullTime   `bun:"deleted_at,nullzero,skipupdate" db:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at,omitempty" yaml:"deleted_at,omitempty" msgpack:"deleted_at,omitempty"`
	DeletedBy sql.NullString `bun:"deleted_by,type:varchar(60),nullzero,skipupdate" db:"deleted_by" json:"deleted_by,omitempty" toml:"deleted_by,omitempty" yaml:"deleted_by,omitempty" msgpack:"deleted_by,omitempty"`
	IsDeleted sql.NullBool   `bun:"is_deleted,nullzero,skipupdate" db:"is_deleted" json:"is_deleted,omitempty" toml:"is_deleted,omitempty" yaml:"is_deleted,omitempty" msgpack:"is_deleted,omitempty"`
}

func (m *MixinModel) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	handleTSZ(m, query)
	handleGUID(m, query)
	return nil
}

// GUIDGen overrides the global generator for every model
func (m *MixinModel) GUIDGen(g Generator) {
	m.gg = g
}

// GUIDGen sets global ID generator for all models
func GUIDGen(g Generator) {
	if g != nil {
		guidMux.Lock()
		gg = g
		guidMux.Unlock()
	} else {
		log.Println("parameter g is nil, ignore")
	}
}

func handleTSZ(m *MixinModel, query bun.Query) {
	if m != nil {
		switch query.Operation() {
		case "INSERT":
			m.CreatedAt = time.Now()
			m.UpdatedAt = m.CreatedAt
		case "UPDATE":
			m.UpdatedAt = time.Now()
		case "DELETE":
			m.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
			m.IsDeleted = sql.NullBool{Bool: true, Valid: true}
		}
	}
}

func handleGUID(m *MixinModel, query bun.Query) {
	if o := query.Operation(); o == "INSERT" && m != nil && len(m.GUID) == 0 {
		var xg = gg
		if m.gg != nil {
			xg = m.gg
		}
		m.GUID = xg()
	}
}
