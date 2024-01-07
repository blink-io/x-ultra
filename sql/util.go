package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
)

type Checker interface {
	IsMySQL() bool
	IsPostgres() bool
	IsSQLite() bool
}

type HealthChecker interface {
	HealthCheck(context.Context) error
}

// IsNoRows .
func IsNoRows(e error) bool {
	return errors.Is(e, sql.ErrNoRows)
}

func doPingFunc(ctx context.Context, f func(context.Context) error) error {
	if err := f(ctx); err != nil && !errors.Is(err, driver.ErrSkip) {
		return err
	}
	return nil
}
