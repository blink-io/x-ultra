package dialect

import (
	"fmt"
	"strings"
	"time"
)

type postgres struct{}

func NewPostgres() Dialect {
	return &postgres{}
}

func (d *postgres) QuoteIdent(s string) string {
	return quoteIdent(s, `"`)
}

func (d *postgres) EncodeString(s string) string {
	// http://www.postgresql.org/docs/9.2/static/sql-syntax-lexical.html
	return `'` + strings.Replace(s, `'`, `''`, -1) + `'`
}

func (d *postgres) EncodeBool(b bool) string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}

func (d *postgres) EncodeTime(t time.Time) string {
	return `'` + t.Format(time.RFC3339Nano) + `'`
}

func (d *postgres) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`E'\\x%x'`, b)
}

func (d *postgres) Placeholder(n int) string {
	return fmt.Sprintf("$%d", n+1)
}
