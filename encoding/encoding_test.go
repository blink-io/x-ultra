package encoding

import (
	"fmt"
	"testing"

	"github.com/apache/incubator-fury/go/fury"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/require"
	"github.com/vmihailenco/msgpack/v5"
)

type Extra struct {
	Info    string
	Enabled bool
	Samples []string
}

type Pyd struct {
	Level int
	Name  string
	Point float64
	Data  []byte
	Extra *Extra
}

func TestEncoding(t *testing.T) {
	fr := fury.NewFury(true)
	require.NoError(t, fr.RegisterTagType("example.Pyd", Pyd{}))
	require.NoError(t, fr.RegisterTagType("example.Extra", Extra{}))

	p1 := &Pyd{
		Level: 999,
		Name:  gofakeit.Name(),
		Point: gofakeit.Float64(),
		Data:  gofakeit.ImageJpeg(60, 60),
		Extra: &Extra{
			Info:    gofakeit.LastName(),
			Enabled: true,
			Samples: []string{
				"zero",
				"one",
				"two",
				"three",
				"four",
				"five",
				"six",
				"seven",
				"eight",
				"nigh",
				"ten",
			},
		},
	}

	type mfn func(interface{}) ([]byte, error)

	mms := map[string]mfn{
		"msgpack": msgpack.Marshal,
		"fury":    fr.Marshal,
		"cbor":    cbor.Marshal,
	}
	require.NotNil(t, mms)

	for n, m := range mms {
		d, err := m(p1)
		require.NoErrorf(t, err, n)
		require.NotNil(t, d)
		fmt.Println("======> Name: ", n, "Len: ", len(d), " <======")
	}
}
