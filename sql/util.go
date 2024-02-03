package sql

import (
	"context"
	"database/sql/driver"
	"errors"
)

// DoPingContext does invoke ping(context.Context).
// Ignore driver.ErrSkip when the Conn does not implement driver.Pinger interface.
func DoPingContext(ctx context.Context, pinger interface {
	PingContext(context.Context) error
}) error {
	if err := pinger.PingContext(ctx); err != nil && !errors.Is(err, driver.ErrSkip) {
		return err
	}
	return nil
}
