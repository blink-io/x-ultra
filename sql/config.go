package sql

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	"github.com/blink-io/x/sql/driver/hooks"
	"go.opentelemetry.io/otel/attribute"
)

var (
	ErrNilConfig = errors.New("config is nil")
)

type Config struct {
	Context         context.Context
	Dialect         string
	Network         string
	Host            string
	Port            int
	Name            string
	User            string
	Password        string
	TLSConfig       *tls.Config
	Params          map[string]string
	DialTimeout     time.Duration
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	ConnInitSQL     string
	ValidationSQL   string
	DriverHooks     []hooks.Hooks
	Loc             *time.Location
	Debug           bool
	Collation       string
	ClientName      string
	Logger          func(format string, args ...any)
	Accessor        string

	// OpenTelemetry
	WithOTel  bool
	OTelAttrs []attribute.KeyValue

	// Additional options
	Additions map[string]string
	dsn       string
}

func SetupConfig(c *Config) *Config {
	if c == nil {
		c = new(Config)
	}
	c.SetDefaults()
	return c
}

func (c *Config) SetDefaults() {
	if c == nil {
		return
	}

	if c.Context == nil {
		c.Context = context.Background()
	}
	if len(c.Network) == 0 {
		c.Network = "tcp"
	}
	if c.Loc == nil {
		c.Loc = time.Local
	}
	if len(c.Dialect) > 0 {
		c.Dialect = GetFormalDialect(c.Dialect)
	}
}

func (c *Config) Validate(ctx context.Context) error {
	if c == nil {
		return ErrNilConfig
	}
	d, ok := IsCompatibleDialect(c.Dialect)
	if !ok {
		return ErrUnsupportedDialect
	}
	switch d {
	case DialectPostgres:
		return ValidatePostgresConfig(c)
	case DialectMySQL:
		return ValidateMySQLConfig(c)
	case DialectSQLite:
		return ValidateSQLiteConfig(c)
	default:
		return ErrUnsupportedDialect
	}
}

func (c *Config) DBInfo() DBInfo {
	return NewDBInfo(c)
}

func (c *Config) DSN() string {
	return c.dsn
}
