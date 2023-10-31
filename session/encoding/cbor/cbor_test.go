package cbor

import (
	"fmt"
	"testing"
	"time"

	"github.com/blink-io/x/session/encoding"
	"github.com/stretchr/testify/require"
)

func TestJSON_1(t *testing.T) {
	p1 := &encoding.Payload{
		Deadline: time.Now().Add(3 * time.Hour),
		Values: map[string]any{
			"name":    "Heison",
			"level":   10,
			"score":   66.7,
			"enabled": true,
			"samples": []string{
				"11", "22", "33",
			},
			"ratios": []float32{
				11.1, 22.2, 33.3, 44.4,
			},
		},
	}

	enc := New()

	b1, err1 := enc.Encode(p1.Deadline, p1.Values)
	require.NoError(t, err1)

	fmt.Println("cbor:   ", string(b1))

	d2, v2, err2 := enc.Decode(b1)
	require.NoError(t, err2)
	fmt.Println("Len:   ", len(b1), "  Deadline:   ", d2, "  Values:   ", v2)
}
