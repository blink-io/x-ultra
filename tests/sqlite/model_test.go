package sqlite

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/blink-io/x/bun/model/mixin"

	"github.com/sanity-io/litter"
	"github.com/uptrace/bun"
)

// Application represents iOS/Android/Windows/OSX/Linux application
type Application struct {
	bun.BaseModel `bun:"applications,alias:applications" db:"-" json:"-" toml:"-" yaml:"-" msgpack:"-"`
	Status        string         `bun:"status,type:varchar(60),notnull" db:"status" json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" msgpack:"status,omitempty"`
	Type          string         `bun:"type,type:varchar(60),notnull" db:"type" json:"type,omitempty" toml:"type,omitempty" yaml:"type,omitempty" msgpack:"type,omitempty"`
	Name          string         `bun:"name,type:varchar(200),notnull" db:"name" json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty" msgpack:"name,omitempty"`
	Code          string         `bun:"code,type:varchar(60),unique,notnull" db:"code" json:"code,omitempty" toml:"code,omitempty" yaml:"code,omitempty" msgpack:"code,omitempty"`
	Description   sql.NullString `bun:"description,type:text" db:"description" json:"description,omitempty" toml:"description,omitempty" yaml:"description,omitempty" msgpack:"description,omitempty"`
	mixin.Model
	mixin.IIDField
}

func (Application) TableName() string {
	return "applications"
}

func (Application) Table() string {
	return "applications"
}

func (Application) Columns() []string {
	return appColumns
}

var appColumns = []string{
	"id",
	"iid",
	"status",
	"name",
	"code",
	"type",
	"created_at",
	"updated_at",
	"deleted_at",
}

func printApp(r *Application) string {
	return litter.Sdump(r)
}

func TestPrintApp(t *testing.T) {
	r := newRandomRecordForApp("test")
	fmt.Println("App: ", printApp(r))
}

func ToAnySlice[T any](a []T) []any {
	t := make([]any, 0, len(a))
	for _, v := range a {
		t = append(t, v)
	}
	return t
}
