package f

import (
	"github.com/keegancsmith/sqlf"
)

type Query = sqlf.Query

type F struct {
}

func New() F {
	return F{}
}

func (F) Sprintf(format string, args ...any) *Query {
	return sqlf.Sprintf(format, args...)
}

func (F) Join(queries []*Query, sep string) *Query {
	return sqlf.Join(queries, sep)
}
