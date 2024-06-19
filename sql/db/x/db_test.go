package x

import (
	"github.com/stretchr/testify/require"
	"testing"

	rdb "github.com/blink-io/x/sql/db"
)

func TestGenDB_1(t *testing.T) {

}

func TestSelectOptions(t *testing.T) {
	ops := []SelectOption{
		SelectApplyQuery(func(q *rdb.SelectQuery) *rdb.SelectQuery {
			return q
		}),
	}
	require.NotNil(t, ops)
}
