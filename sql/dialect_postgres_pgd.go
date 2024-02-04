//go:build pgdriver

package sql

import (
	"context"
	"database/sql/driver"

	"github.com/uptrace/bun/driver/pgdriver"
)

func GetPostgresPgdConnector(ctx context.Context, c *Config) (driver.Connector, error) {
	addr := hostportToAddr(c.Host, c.Port)
	pgops := []pgdriver.Option{
		pgdriver.WithApplicationName(c.ClientName),
		pgdriver.WithNetwork(c.Transport),
		pgdriver.WithAddr(addr),
		pgdriver.WithDatabase(c.Name),
		pgdriver.WithUser(c.User),
		pgdriver.WithPassword(c.Password),
		pgdriver.WithInsecure(c.TLSConfig == nil),
		pgdriver.WithTLSConfig(c.TLSConfig),
	}
	if c.DialTimeout > 0 {
		pgops = append(pgops, pgdriver.WithDialTimeout(c.DialTimeout))
	}
	cc := pgdriver.NewConnector(pgops...)
	return cc, nil
}

func getRawPostgresPgdDriver() driver.Driver {
	// Notes: Unable to invoke &stdlib.Driver{} directly.
	// Because the "configs" field inside the driver is not initialized.
	return pgdriver.NewDriver()
}
