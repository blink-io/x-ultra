package tests

import (
	"testing"

	"github.com/jhump/protoreflect/dynamic"
	"github.com/stretchr/testify/require"
)

func TestErr1(t *testing.T) {
	extgry := dynamic.NewExtensionRegistryWithDefaults()
	require.NotNil(t, extgry)
}
