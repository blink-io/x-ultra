package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xo/dburl"
)

func TestDBURL_1(t *testing.T) {
	urlGen := dburl.GenScheme("sqlite")
	require.NotNil(t, urlGen)
}
