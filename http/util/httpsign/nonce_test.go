package httpsign

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/clock"

	"github.com/stretchr/testify/require"
)

func TestNonceInCache(t *testing.T) {
	clock.Freeze(clock.Now())
	defer clock.Unfreeze()

	// setup
	nc, err := newNonceCache(
		nil,
		1*time.Second,
	)
	if err != nil {
		t.Error("Got unexpected error from newNonceCache:", err)
	}

	beginT := time.Now()
	// nothing in cache, it should be valid
	inCache, err := nc.inCache("0")
	require.NoError(t, err)
	if inCache {
		t.Error("Check should be valid, but failed.")
	}

	// second time around it shouldn't be
	inCache, err = nc.inCache("0")
	require.NoError(t, err)
	if !inCache {
		t.Error("Check should be invalid, but passed.")
	}

	// check some other value
	clock.Advance(999 * clock.Millisecond)
	inCache, err = nc.inCache("1")
	require.NoError(t, err)
	if inCache {
		t.Error("Check should be valid, but failed.", err)
	}

	// age off first value, then it should be valid
	clock.Advance(1 * clock.Millisecond)

	endT := time.Now()
	cost := endT.Sub(beginT)
	fmt.Printf("cost in secs: %d\n", int64(cost.Seconds()))
	fmt.Printf("cost in microsecs: %d\n", int64(cost.Microseconds()))
	fmt.Printf("cost in millsecs: %d\n", int64(cost.Milliseconds()))
	fmt.Printf("cost in nanosecs: %d\n", int64(cost.Nanoseconds()))

	inCache, err = nc.inCache("0")
	require.NoError(t, err)
	if inCache {
		t.Error("Check should be valid, but failed.")
	}
}
