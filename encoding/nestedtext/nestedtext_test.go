package nestedtext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNT_Decode_1(t *testing.T) {
	input := `
# Example for a NestedText dict
a: Hello
b: World
`
	enc := New()
	var dm = make(map[string]any)
	err := enc.Unmarshal([]byte(input), dm)
	require.NoError(t, err)
}

func TestNT_Encode_1(t *testing.T) {
	var config = map[string]any{
		"timeout": 20,
		"ports":   []any{6483, 8020, 9332},
	}
	enc := New()
	data, err := enc.Marshal(config)
	require.NoError(t, err)

	fmt.Println("------------------------------")
	fmt.Println(string(data))
}
