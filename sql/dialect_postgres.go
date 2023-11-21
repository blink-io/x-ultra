package sql

import (
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
	dialectFuncs[dn] = func() schema.Dialect {
		return pgdialect.New()
	}
	dsnFuncs[dn] = PostgresDSN
}

func PostgresDSN(o *Options) string {
	name := o.Name
	host := o.Host
	port := o.Port
	user := o.User
	password := o.Password
	dialTimeout := o.DialTimeout
	tlsConfig := o.TLSConfig
	params := o.Params
	if params == nil {
		params = make(map[string]string)
	}
	if len(o.ClientName) > 0 {
		params[pgparams.APPLICATION_NAME] = o.ClientName
	}
	if len(o.Collation) > 0 {
		params[pgparams.CLIENT_ENCODING] = o.Collation
	}

	pgcc := pgconn.Config{
		Database:      name,
		Host:          host,
		Port:          uint16(port),
		User:          user,
		Password:      password,
		TLSConfig:     tlsConfig,
		RuntimeParams: handlePostgresParams(params),
	}
	if dialTimeout > 0 {
		pgcc.ConnectTimeout = dialTimeout
	}

	cc := &pgx.ConnConfig{
		Config: pgcc,
		Tracer: &tracelog.TraceLog{Logger: pgxzap.NewLogger(zap.L()), LogLevel: tracelog.LogLevelInfo},
	}
	dsn := stdlib.RegisterConnConfig(cc)
	return dsn
}

func handlePostgresParams(params map[string]string) map[string]string {
	newParams := make(map[string]string)
	for k, v := range params {
		newParams[k] = v
	}
	return newParams
}
