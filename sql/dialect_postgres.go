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
	dialectors[dn] = func(ctx context.Context, ops ...DOption) schema.Dialect {
		return pgdialect.New()
	}
	dsnors[dn] = PostgresDSN
}

func PostgresDSN(ctx context.Context, o *Options) (string, error) {
	cc := ToPGXConfig(o)
	dsn := stdlib.RegisterConnConfig(cc)
	return dsn, nil
}

func ToPGXConfig(o *Options) *pgx.ConnConfig {
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
		params[pgparams.ApplicationName] = o.ClientName
	}
	if len(o.Collation) > 0 {
		params[pgparams.ClientEncoding] = o.Collation
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
	if o.Debug {
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
