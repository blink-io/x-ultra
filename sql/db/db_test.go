package db

import (
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun/schema"
	"testing"
)

func TestSafeQuery_1(t *testing.T) {
	q := "abc"
	ss := doSafeQuery(q, "a", "b", "c")
	require.NotNil(t, ss)
}

func doSafeQuery(q string, args ...any) schema.QueryWithArgs {
	ss := SafeQuery(q, args)
	return ss
}
