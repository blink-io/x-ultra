package model

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aarondl/opt/omitnull"
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

	ColumnNames = &columnNames{
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

type columnNames struct {
	// Require columnNames
	ID        string
	CreatedAt string
	UpdatedAt string
	// Optional columnNames
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
	ID        int64     `bun:"id,pk,autoincrement" db:"id" json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty" msgpack:"id,omitempty"`
	GUID      string    `bun:"guid,notnull,unique,type:varchar(60)" db:"guid" json:"guid,omitempty" toml:"guid,omitempty" yaml:"guid,omitempty" msgpack:"guid,omitempty"`
	CreatedAt time.Time `bun:"created_at,notnull,skipupdate" db:"created_at" json:"created_at,omitempty" toml:"created_at,omitempty" yaml:"created_at,omitempty" msgpack:"created_at,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,notnull" db:"updated_at" json:"updated_at,omitempty" toml:"updated_at,omitempty" yaml:"updated_at,omitempty" msgpack:"updated_at,omitempty"`
	// Optional fields for tables
	CreatedBy omitnull.Val[string]    `bun:"created_by,type:varchar(60),nullzero,skipupdate" db:"created_by" json:"created_by,omitempty" toml:"created_by,omitempty" yaml:"created_by,omitempty" msgpack:"created_by,omitempty"`
	UpdatedBy omitnull.Val[string]    `bun:"updated_by,type:varchar(60),nullzero" db:"updated_by" json:"updated_by,omitempty" toml:"updated_by,omitempty" yaml:"updated_by,omitempty" msgpack:"updated_by,omitempty"`
	DeletedAt omitnull.Val[time.Time] `bun:"deleted_at,nullzero,skipupdate" db:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at,omitempty" yaml:"deleted_at,omitempty" msgpack:"deleted_at,omitempty"`
	DeletedBy omitnull.Val[string]    `bun:"deleted_by,type:varchar(60),nullzero,skipupdate" db:"deleted_by" json:"deleted_by,omitempty" toml:"deleted_by,omitempty" yaml:"deleted_by,omitempty" msgpack:"deleted_by,omitempty"`
	IsDeleted omitnull.Val[bool]      `bun:"is_deleted,nullzero,skipupdate" db:"is_deleted" json:"is_deleted,omitempty" toml:"is_deleted,omitempty" yaml:"is_deleted,omitempty" msgpack:"is_deleted,omitempty"`
}

func (m *MixinModel) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	handleTSZ(m, ctx, query)
	handleGUID(m, ctx, query)
	return nil
}

// GUIDGen sets global I generator for all models
func GUIDGen(g Generator) {
	if g != nil {
		guidMux.Lock()
		gg = g
		guidMux.Unlock()
	} else {
		log.Println("parameter g is nil, ignore")
	}
}

func handleTSZ(m *MixinModel, ctx context.Context, query bun.Query) {
	if m != nil {
		switch query.(type) {
		case *bun.InsertQuery:
			m.CreatedAt = time.Now()
			m.UpdatedAt = m.CreatedAt
		case *bun.UpdateQuery:
			m.UpdatedAt = time.Now()
		case *bun.DeleteQuery:
			m.DeletedAt = omitnull.From(time.Now())
			m.IsDeleted = omitnull.From(true)
		}
	}
}

func handleGUID(m *MixinModel, ctx context.Context, query bun.Query) {
	if o := query.Operation(); o == "INSERT" && m != nil && len(m.GUID) == 0 {
		m.GUID = gg()
	}
}
