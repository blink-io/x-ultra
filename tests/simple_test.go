package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type User struct {
	name string
}

func TestSim_1(t *testing.T) {
	cn := func(ok bool) string {
		if ok {
			return "a"
		}
		return "b"
	}
	name := "a"
	name2 := cn(false)

	u1 := new(User)
	u1.name = name

	u2 := new(User)
	u2.name = name2

	require.Equal(t, u1, u2)
}
