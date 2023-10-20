package sqlitedialect

import (
	"time"

	"github.com/uptrace/bun/dialect/sqlitedialect"
)

type Dialect struct {
	*sqlitedialect.Dialect
	opt *options
}

func New(ops ...Option) *Dialect {
	opt := applyOptions(ops...)
	d := &Dialect{Dialect: sqlitedialect.New(), opt: opt}
	d.Features()
	return d
}

// AppendTime in the schema.BaseDialect uses the UTC timezone.
// Let the developers make the decision.
func (d *Dialect) AppendTime(b []byte, tm time.Time) []byte {
	b = append(b, '\'')
	if d.opt.utc {
		tm = tm.UTC()
	}
	b = tm.AppendFormat(b, "2006-01-02 15:04:05.999999-07:00")
	b = append(b, '\'')
	return b
}
