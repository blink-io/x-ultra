package pg

import (
	"github.com/aarondl/opt/omitnull"
	"time"

	"github.com/blink-io/x/id"
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
	r.CreatedBy = omitnull.From(gofakeit.Name())
	r.UpdatedBy = omitnull.From(gofakeit.Name())
	r.DeletedAt = omitnull.From(tnow)
	return r
}
