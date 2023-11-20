package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ulikunitz/xz"
)

func TestZstd_1(t *testing.T) {
	fname := "../testdata/littlefs.tar.gz"

	fl, err := os.Open(fname)
	require.NoError(t, err)

	w, err := xz.NewWriter(fl)
	require.NoError(t, err)

	defer w.Close()
}
