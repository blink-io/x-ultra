//go:build pq

package sql

import (
	"context"
	"database/sql/driver"

	"github.com/lib/pq"
)

func GetPostgresPgConnector(ctx context.Context, c *Config) (driver.Connector, error) {
	//addr := hostportToAddr(c.Host, c.Port)
	url := ""
	dsn, err := pq.ParseURL(url)
	cc, err := pq.NewConnector(dsn)
	return cc, err
}

func getRawPostgresPqDriver() driver.Driver {
	return &pq.Driver{}
}
