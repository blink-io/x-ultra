package sql

import (
	"context"
	"database/sql/driver"

	pgparams "github.com/blink-io/x/postgres/params"
	"github.com/life4/genesis/slices"

	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"
)

var compatiblePostgresDialects = []string{
	DialectPostgres,
	"postgresql",
	"pg",
	"pgx",
}

func init() {
	dn := DialectPostgres
	drivers[dn] = GetPostgresDriver
	dsners[dn] = GetPostgresDSN
}

func GetPostgresDSN(dialect string) (Dsner, error) {
	if !IsCompatiblePostgresDialect(dialect) {
		return nil, ErrUnsupportedDialect
	}
	return func(ctx context.Context, c *Config) (string, error) {
		cc := ToPGXConfig(c)
		dsn := stdlib.RegisterConnConfig(cc)
		return dsn, nil
	}, nil
}

func ToPGXConfig(c *Config) *pgx.ConnConfig {
	name := c.Name
	host := c.Host
	port := c.Port
	user := c.User
	password := c.Password
	dialTimeout := c.DialTimeout
	tlsConfig := c.TLSConfig
	params := c.Params
	if params == nil {
		params = make(map[string]string)
	}
	if len(c.ClientName) > 0 {
		params[pgparams.ApplicationName] = c.ClientName
	}
	if len(c.Collation) > 0 {
		params[pgparams.ClientEncoding] = c.Collation
	}

	pgcc, err := pgconn.ParseConfig("")
	if err != nil {
		// This can be happened
		panic(err)
	}

	pgcc.Database = name
	pgcc.Host = host
	pgcc.Port = uint16(port)
	pgcc.User = user
	pgcc.Password = password
	pgcc.TLSConfig = tlsConfig
	pgcc.RuntimeParams = handlePostgresParams(params)
	if dialTimeout > 0 {
		pgcc.ConnectTimeout = dialTimeout
	}

	cc, err := pgx.ParseConfig("")
	if err != nil {
		// This can't be happened
		panic(err)
	}
	cc.Config = *pgcc
	traceLogLevel := tracelog.LogLevelInfo
	if c.Debug {
		traceLogLevel = tracelog.LogLevelDebug
	}
	cc.Tracer = &tracelog.TraceLog{Logger: pgxzap.NewLogger(zap.L()), LogLevel: traceLogLevel}
	return cc
}

func handlePostgresParams(params map[string]string) map[string]string {
	newParams := make(map[string]string)
	for k, v := range params {
		newParams[k] = v
	}
	return newParams
}

func IsCompatiblePostgresDialect(dialect string) bool {
	i := slices.FindIndex(compatiblePostgresDialects, func(i string) bool {
		return i == dialect
	})
	return i > -1
}

func GetPostgresDriver(dialect string) driver.Driver {
	if IsCompatiblePostgresDialect(dialect) {
		return stdlib.GetDefaultDriver()
	}
	return nil
}
