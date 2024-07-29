package sqlite

import (
	"github.com/aarondl/opt/omitnull"
	"time"

	"github.com/google/uuid"

	"github.com/blink-io/x/id"
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

func newRandomRecordForUserDevice(from string) *UserDevice {
	now := time.Now().Local()
	r := new(UserDevice)
	r.GUID = id.UUID()
	r.Name = "from-" + from + "-" + gofakeit.Name()
	r.CreatedAt = now
	r.UpdatedAt = now
	return r
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
	r.CreatedBy = omitnull.From(gofakeit.Name())
	r.UpdatedBy = omitnull.From(gofakeit.Name())
	r.DeletedAt = omitnull.From(tnow)
	return r
}

func appModel() *Application {
	return new(Application)
}
