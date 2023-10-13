package sql

import (
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
	options := o.Options

	//debug := o.Debug

	cc := &pgx.ConnConfig{
		Config: pgconn.Config{
			Database:       name,
			Host:           host,
			Port:           uint16(port),
			User:           user,
			Password:       password,
			TLSConfig:      tlsConfig,
			ConnectTimeout: dialTimeout,
			// TODO Do we need to check them?
			RuntimeParams: options,
		},
	}
	cc.Tracer = &tracelog.TraceLog{Logger: pgxzap.NewLogger(zap.L()), LogLevel: tracelog.LogLevelInfo}
	dsn := stdlib.RegisterConnConfig(cc)
	return dsn
}
