package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vmihailenco/tagparser/v2"
)

func TestTag_1(t *testing.T) {
	tag := tagparser.Parse("some_name,key:value,key2:'complex value',intv:1223")
	require.NotNil(t, tag)
}
