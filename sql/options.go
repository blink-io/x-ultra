package sql

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/blink-io/x/sql/hooks"
	"go.opentelemetry.io/otel/attribute"

	"github.com/uptrace/bun"
)

type Options struct {
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
	ConnInitSQL     string
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	ValidationSQL   string
	DriverHooks     []hooks.Hooks
	QueryHooks      []bun.QueryHook
	Loc             *time.Location
	Debug           bool
	WithOTel        bool
	Collation       string
	ClientName      string
	Attrs           []attribute.KeyValue
	dsn             string
}

func setupOptions(o *Options) *Options {
	if o == nil {
		o = new(Options)
	}
	o.SetDefaults()
	return o
}

func (o *Options) SetDefaults() {
	if o == nil {
		return
	}
	if len(o.Network) == 0 {
		o.Network = "tcp"
	}
	if o.Loc == nil {
		o.Loc = time.Local
	}
}

func (o *Options) Validate() error {
	return nil
}
