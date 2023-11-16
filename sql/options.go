package sql

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	"github.com/blink-io/x/sql/hooks"

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
	Options         map[string]string
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
	WithOTEL        bool
	Collation       string
	dsn             string
}

func setupOptions(o *Options) (*Options, error) {
	if o == nil {
		return nil, errors.New("idb config cannot be empty")
	}
	if len(o.Network) == 0 {
		o.Network = "tcp"
	}
	if o.Loc == nil {
		o.Loc = time.Local
	}
	return o, nil
}
