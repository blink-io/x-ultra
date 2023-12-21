package xsql_test

import (
	"fmt"
	"log/slog"

	xsql "github.com/blink-io/x/sql"
	"github.com/doug-martin/goqu/v9"
)

type logger func(format string, args ...any)

func (l logger) Printf(format string, args ...any) {
	l(format, args...)
}

var _ goqu.Logger = (logger)(nil)

func handleDBQ(db *xsql.DBQ) {
	db.Logger(logger(func(format string, args ...any) {
		slog.Default().Info(fmt.Sprintf(format, args...))
	}))
}

func handleDBX(db *xsql.DBX) {
	db.LogFunc = func(format string, args ...interface{}) {
		slog.Default().Info(fmt.Sprintf(format, args...))
	}
}
