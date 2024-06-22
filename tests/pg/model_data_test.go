package pg

import (
	"time"

	"github.com/blink-io/x/id"
	xsql "github.com/blink-io/x/sql"
	"github.com/brianvoe/gofakeit/v6"
)

func newRandomRecordForApp(from string) *Application {
	tnow := time.Now().Local()
	r := new(Application)
	r.GUID = id.ShortUUID()
	r.Name = "from-" + from + "-" + gofakeit.Name()
	r.Code = "code-" + gofakeit.Name()
	r.Type = "type-001"
	r.Status = "ok"
	r.CreatedAt = tnow
	r.UpdatedAt = tnow
	r.CreatedBy = xsql.ValidNull[string](gofakeit.Name())
	r.UpdatedBy = xsql.ValidNull[string](gofakeit.Name())
	r.DeletedAt = xsql.ValidNull[time.Time](tnow)
	return r
}
