package sql

import (
	"context"
	"database/sql/driver"

	"modernc.org/sqlite"
)

var compatibleSQLiteDialects = []string{
	DialectSQLite,
	"sqlite3",
}

func init() {
	d := DialectSQLite
	//drivers[dn] = GetSQLiteDriver
	//dsners[dn] = GetSQLiteDSN
	connectors[d] = GetQLiteConnector
}

type SQLiteOptions struct {
}

func GetSQLiteDSN(dialect string) (Dsner, error) {
	if !IsCompatibleSQLiteDialect(dialect) {
		return nil, ErrUnsupportedDialect
	}
	return func(ctx context.Context, c *Config) (string, error) {
		dsn := c.Host
		return dsn, nil
	}, nil
}

func IsCompatibleSQLiteDialect(dialect string) bool {
	return isCompatibleDialect(dialect, compatibleSQLiteDialects)
}

func GetSQLiteDriver(dialect string) (driver.Driver, error) {
	if IsCompatibleSQLiteDialect(dialect) {
		return &sqlite.Driver{}, nil
	}
	return nil, ErrUnsupportedDriver
}

func GetQLiteConnector(ctx context.Context, c *Config) (driver.Connector, error) {
	dsn := toSQLiteDSN(c)
	drv := wrapDriverHooks(getRawSQLiteDriver(), c.DriverHooks...)
	return &dsnConnector{dsn: dsn, driver: drv}, nil
}

func AdditionsToSQLiteOptions(adds map[string]string) *SQLiteOptions {
	opts := new(SQLiteOptions)
	return opts
}

func toSQLiteDSN(c *Config) string {
	dsn := c.Host
	return dsn
}

func getRawSQLiteDriver() *sqlite.Driver {
	return &sqlite.Driver{}
}
