package x

import (
	"testing"

	rdb "github.com/blink-io/x/bun"
	"github.com/stretchr/testify/require"
)

func TestGenDB_1(t *testing.T) {

}

func TestSelectOptions(t *testing.T) {
	ops := []SelectOption{
		WithSelectQuery(func(q *rdb.SelectQuery) *rdb.SelectQuery {
			return q
		}),
	}
	require.NotNil(t, ops)
}