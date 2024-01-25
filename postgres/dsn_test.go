package postgres

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDSN_1(t *testing.T) {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=test sslmode=disable"
	ss, err := ParseDSN(dsn)
	require.NoError(t, err)
	require.NotNil(t, ss)

	fmt.Println("done")

	ndsn, err := ss.ToDSN()
	require.NoError(t, err)
	require.NotNil(t, ndsn)

	fmt.Println("Reverted DSN: ", ndsn)

}

func TestParseURL_1(t *testing.T) {
	urlstr := "postgresql://user:pass@localhost/mydatabase/?sslmode=disable"
	ss, err := ParseURL(urlstr)
	require.NoError(t, err)
	require.NotNil(t, ss)

	fmt.Println("done")
}

func TestParseURL_2(t *testing.T) {
	urlstr := "postgresql://user:pass@localhost/mydatabase/?sslmode=disable"

	ss := NewSettings()
	err := ss.ParseURL(urlstr)
	require.NoError(t, err)

	fmt.Println("done")
}
