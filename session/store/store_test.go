package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNilStruct(t *testing.T) {
	ns := NilStruct

	require.Nilf(t, ns, "NilStruct is really nil")
}

func TestTokenMap_1(t *testing.T) {
	var tm = NewTokenMap()
	tm.SetNil("abc")
	tm.SetNil("efg")

	fmt.Println("done")
}
