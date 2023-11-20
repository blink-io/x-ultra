package dsn

import (
	"net/url"
)

type SQLite struct {
}

var _ Processor = (*SQLite)(nil)

func (m *SQLite) Parse(dsn string) (*url.URL, error) {
	u := &url.URL{
		Scheme: m.Name(),
	}
	return u, nil
}

func (m *SQLite) Convert(url *url.URL) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *SQLite) Name() string {
	return "sqlite"
}

func (m *SQLite) Schemes() []string {
	return []string{
		"sqlite",
		"sqlite3",
	}
}
