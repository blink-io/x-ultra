package dsn

import (
	"net/url"
)

type MySQL struct {
}

var _ Processor = (*MySQL)(nil)

func (m *MySQL) Parse(dsn string) (*url.URL, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) Convert(url *url.URL) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) Name() string {
	return "mysql"
}

func (m *MySQL) Schemes() []string {
	return []string{
		"mysql",
	}
}
