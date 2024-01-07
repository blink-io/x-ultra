package sql

import (
	"context"

	pgparams "github.com/blink-io/x/postgres/params"

	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/schema"
	"go.uber.org/zap"
)

const (
	DialectPostgres = "postgres"
)

func init() {
	dn := DialectPostgres
	drivers[dn] = stdlib.GetDefaultDriver()
	dialectors[dn] = NewPostgresDialect
	dsners[dn] = PostgresDSN
}

func NewPostgresDialect(ctx context.Context, ops ...DialectOption) schema.Dialect {
	return pgdialect.New()
}

func PostgresDSN(ctx context.Context, c *Config) (string, error) {
	cc := ToPGXConfig(c)
	dsn := stdlib.RegisterConnConfig(cc)
	return dsn, nil
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
