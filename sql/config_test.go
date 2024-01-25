package sql

import (
	"context"
	"fmt"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
	"github.com/xo/dburl"
)

func TestOptions_1(t *testing.T) {
	var o *Config
	o = SetupConfig(o)
	require.NotNil(t, o)
}

func TestOptions_Validate(t *testing.T) {
	var o *Config
	o = SetupConfig(o)
	verr := o.Validate(context.Background())
	require.NoError(t, verr)
}

func TestParseURL_1(t *testing.T) {
	urlstr := "postgresql://user:pass@localhost/mydatabase/?sslmode=disable"
	c, err := ParseURL(urlstr)
	require.NoError(t, err)
	require.NotNil(t, c)

	du, err := dburl.Parse(urlstr)
	require.NoError(t, err)
	require.NotNil(t, du)

	fmt.Println("EscapedPath: ", du.EscapedPath())

	litter.Dump(du)

	fmt.Println()

	litter.Dump(du.Query())
}
