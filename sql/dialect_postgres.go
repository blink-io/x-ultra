package sql

import (
	"context"
	"database/sql/driver"
	"log/slog"

	"github.com/blink-io/x/cast"
	pgparams "github.com/blink-io/x/postgres/params"
	pgxslog "github.com/blink-io/x/postgres/pgx/logger/slog"
	pgxotel "github.com/blink-io/x/postgres/pgx/tracer/otel"
	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"go.uber.org/zap"
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

func ValidatePostgresConfig(c *Config) error {
	return nil
}

func GetPostgresDSN(dialect string) (Dsner, error) {
	if !IsCompatiblePostgresDialect(dialect) {
		return nil, ErrUnsupportedDialect
	}
	return func(ctx context.Context, c *Config) (string, error) {
		cc, _, err := ToPGXConfig(c)
		if err != nil {
			return "", err
		}
		dsn := stdlib.RegisterConnConfig(cc)
		return dsn, nil
	}, nil
}

func GetPostgresDriver(dialect string) (driver.Driver, error) {
	if IsCompatiblePostgresDialect(dialect) {
		return getRawPostgresDriver(), nil
	}
	return nil, ErrUnsupportedDriver
}

func GetPostgresConnector(ctx context.Context, c *Config) (driver.Connector, error) {
	cc, aopt, err := ToPGXConfig(c)
	if err != nil {
		return nil, err
	}
	if aopt.usePool {
		ppc, errv := pgxpool.ParseConfig("")
		if errv != nil {
			return nil, errv
		}
		ppc.ConnConfig = cc
		pool, errp := pgxpool.NewWithConfig(ctx, ppc)
		if errp != nil {
			return nil, errp
		}
		c.dsn = ppc.ConnString()
		return stdlib.GetPoolConnector(pool), nil
	} else {
		c.dsn = stdlib.RegisterConnConfig(cc)
		drv := wrapDriverHooks(getRawPostgresDriver(), c.DriverHooks...)
		return &dsnConnector{dsn: c.dsn, driver: drv}, nil
	}
}

func ToPGXConfig(c *Config) (*pgx.ConnConfig, *PostgresOptions, error) {
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
		return nil, nil, err
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
		return nil, nil, err
	}

	aopts := AdditionsToPostgresOptions(c.Additions)

	cc.Config = *pgcc
	traceLogLevel := tracelog.LogLevelInfo
	if c.Debug {
		traceLogLevel = tracelog.LogLevelDebug
	}

	if aopts.trace == PostgresTraceOTel {
		cc.Tracer = pgxotel.NewTracer()
	} else {
		var tlogger tracelog.Logger
		if l := c.Logger; l != nil {
			tlogger = tracelog.LoggerFunc(func(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
				alen := len(data) * 2
				args := make([]any, alen+1)
				var i = 0
				for k, v := range data {
					args[i] = k
					args[i+1] = v
					i = i + 2
				}
				l("msg: %s, data: %#v", args...)
			})
		} else {
			if aopts.tracelog == PostgresTracelogZap {
				tlogger = pgxzap.NewLogger(zap.L())
			} else {
				tlogger = pgxslog.NewLogger(slog.Default())
			}
		}
		cc.Tracer = &tracelog.TraceLog{Logger: tlogger, LogLevel: traceLogLevel}
	}
	return cc, aopts, nil
}

func ToPGXPoolConfig(c *Config) (*pgxpool.Config, *PostgresOptions, error) {
	ppc, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, nil, err
	}

	pgcc, aopt, err := ToPGXConfig(c)
	if err != nil {
		return nil, nil, err
	}

	ppc.ConnConfig = pgcc
	return ppc, aopt, nil
}

func AdditionsToPostgresOptions(adds map[string]string) *PostgresOptions {
	opts := new(PostgresOptions)
	if adds != nil {
		opts.trace = adds[AdditionTrace]
		opts.tracelog = adds[AdditionTracelog]
		opts.usePool = cast.ToBool(adds[AdditionUsePool])
	}
	return opts
}

func getRawPostgresDriver() driver.Driver {
	// Notes: Unable to invoke &stdlib.Driver{} directly.
	// Because the "configs" field inside the driver is not initialized.
	return stdlib.GetDefaultDriver()
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
