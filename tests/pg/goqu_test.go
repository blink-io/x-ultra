package pg

import (
	"fmt"
	"log/slog"

	"github.com/blink-io/x/sql/dbq"
	xsql "github.com/blink-io/x/sql/dbx"
	"github.com/doug-martin/goqu/v9"
)

type logger func(format string, args ...any)

func (l logger) Printf(format string, args ...any) {
	l(format, args...)
}

var _ goqu.Logger = (logger)(nil)

func handleDBQ(db *dbq.DB) {
	db.Logger(logger(func(format string, args ...any) {
		slog.Default().Info(fmt.Sprintf(format, args...))
	}))
}

func handleDBX(db *xsql.DB) {

}
