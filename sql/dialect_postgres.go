package sql

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/schema"
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
			//AfterConnect: func(ctx context.Context, conn *pgconn.PgConn) error {
			//	log.Infof("PostgreSQL database is connected")
			//	return nil
			//},
		},
	}
	dsn := stdlib.RegisterConnConfig(cc)
	return dsn
}
