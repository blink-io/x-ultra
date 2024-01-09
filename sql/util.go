package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
)

// IsNoRows .
func IsNoRows(e error) bool {
	return errors.Is(e, sql.ErrNoRows)
}

func DoPingContext(ctx context.Context, db *sql.DB) error {
	if err := db.PingContext(ctx); err != nil && !errors.Is(err, driver.ErrSkip) {
		return err
	}
	return nil
}
