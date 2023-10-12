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
	ConnMaxLifetime *time.Duration
	ConnMaxIdleTime *time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	ValidationSQL   string
	SQLHooks        []hooks.Hooks
	QueryHooks      []bun.QueryHook
	UseOtel         bool
	Loc             *time.Location
	Debug           bool
	dsn             string
}

func setupOptions(c *Options) (*Options, error) {
	if c == nil {
		return nil, errors.New("db config cannot be empty")
	}
	if len(c.Network) == 0 {
		c.Network = "tcp"
	}
	return c, nil
}
