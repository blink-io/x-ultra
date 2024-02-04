package sql

import (
	"database/sql/driver"
)

const (
	PostgresTraceOTel = "otel"

	PostgresTracelogZap = "zap"
)

var compatiblePostgresDialects = []string{
	DialectPostgres,
	"postgresql",
	"pg",
	"pgx",
}

func init() {
	d := DialectPostgres
	//drivers[dn] = GetPostgresDriver
	//dsners[dn] = GetPostgresDSN
	connectors[d] = GetPostgresConnector
}

type PostgresOptions struct {
	trace    string
	tracelog string
	usePool  bool
}

func GetPostgresDriver(dialect string) (driver.Driver, error) {
	if IsCompatiblePostgresDialect(dialect) {
		return getRawPostgresDriver(), nil
	}
	return nil, ErrUnsupportedDriver
}

func ValidatePostgresConfig(c *Config) error {
	return nil
}

func handlePostgresParams(params map[string]string) map[string]string {
	newParams := make(map[string]string)
	for k, v := range params {
		newParams[k] = v
	}
	return newParams
}

func IsCompatiblePostgresDialect(dialect string) bool {
	return isCompatibleDialectIn(dialect, compatiblePostgresDialects)
}
