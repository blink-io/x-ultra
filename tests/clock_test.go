package tests

import (
	"testing"

	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"
)

func TestClock_1(t *testing.T) {
	clk := clockwork.NewRealClock()
	require.NotNil(t, clk)
}
