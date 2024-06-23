package sqlite

import (
	"time"

	"github.com/google/uuid"

	"github.com/blink-io/x/id"
	xsql "github.com/blink-io/x/sql"
	"github.com/brianvoe/gofakeit/v6"
)

func newRandomUserMap() map[string]any {
	values := map[string]any{
		"guid":       uuid.NewString(),
		"username":   gofakeit.Name(),
		"location":   gofakeit.City(),
		"level":      gofakeit.Int8(),
		"profile":    gofakeit.AppName(),
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	return values
}

func newRandomRecordForApp(from string) *Application {
	tnow := time.Now().Local()
	r := new(Application)
	r.GUID = id.UUID()
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

func appModel() *Application {
	return new(Application)
}
