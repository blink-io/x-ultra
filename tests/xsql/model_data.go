package bun

import (
	"time"

	"github.com/blink-io/x/id"
	xsql "github.com/blink-io/x/sql"
	"github.com/brianvoe/gofakeit/v6"
)

func newRandomRecordForApp(from string) *Application {
	tnow := time.Now().Local()
	r := new(Application)
	r.ID = id.ShortUUID()
	r.Name = "from-" + from + "-" + gofakeit.Name()
	r.Code = "code-" + gofakeit.Name()
	r.Type = "type-001"
	r.Status = "ok"
	r.IID = int64(gofakeit.Int32())
	r.CreatedAt = tnow
	r.UpdatedAt = tnow
	r.CreatedBy = xsql.NullString(gofakeit.Name())
	r.UpdatedBy = xsql.NullString(gofakeit.Name())
	r.DeletedAt = xsql.NullTime(tnow)
	return r
}
