package dsn

import (
	"net/url"
)

type Postgres struct {
}

var _ Processor = (*Postgres)(nil)

func (m *Postgres) Parse(dsn string) (*url.URL, error) {
	u := &url.URL{
		Scheme: m.Name(),
	}
	return u, nil
}

func (m *Postgres) Convert(url *url.URL) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Postgres) Name() string {
	return "postgres"
}

func (m *Postgres) Schemes() []string {
	return []string{
		"postgres",
		"postgresql",
		"pg",
	}
}
