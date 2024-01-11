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
	if loc := d.opt.loc; loc != nil {
		tm = tm.In(loc)
	}
	b = tm.AppendFormat(b, time.RFC3339Nano)
	b = append(b, '\'')
	return b
}
