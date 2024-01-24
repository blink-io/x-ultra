package sql

import (
	"context"

	"github.com/blink-io/x/cast"
	"github.com/xo/dburl"
)

func ParseURL(url string) (*Config, error) {
	u, err := dburl.Parse(url)
	if err != nil {
		return nil, err
	}

	pwd, _ := u.User.Password()
	c := &Config{
		Context:  context.Background(),
		Network:  u.Transport,
		Dialect:  u.Driver,
		Host:     u.Host,
		Port:     cast.ToInt(u.Port()),
		User:     u.User.Username(),
		Password: pwd,
		Name:     u.Path,
		dsn:      u.DSN,
	}
	return c, nil
}
