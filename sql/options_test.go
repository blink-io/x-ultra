package sql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptions_1(t *testing.T) {
	var o *Options
	o = setupOptions(o)
	require.NotNil(t, o)
}

func TestOptions_Validate(t *testing.T) {
	var o *Options
	o = setupOptions(o)
	verr := o.Validate()
	require.NoError(t, verr)
}
