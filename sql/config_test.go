package sql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptions_1(t *testing.T) {
	var o *Config
	o = setupConfig(o)
	require.NotNil(t, o)
}

func TestOptions_Validate(t *testing.T) {
	var o *Config
	o = setupConfig(o)
	verr := o.Validate(context.Background())
	require.NoError(t, verr)
}
