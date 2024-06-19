package hyper

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

type testStruct struct {
	bun.BaseModel `bun:"table:test_struct,alias:rt"`

	ID int `bun:"id,pk,autoincrement"`
}

func TestHyperbunIDType(t *testing.T) {
	var db = NewContext(context.Background(), nil)
	r, err := ByID[testStruct, int64](db, 1)
	if err != nil {
		return
	}
	require.NotNil(t, r)
}

func TestHyperbunTableForType(t *testing.T) {
	assert.Equal(t, "test_struct", hyperbunTableForType[testStruct]())
}

func TestAnnotateEven(t *testing.T) {
	assert.Equal(t,
		"performing TestAnnotate hello='world' id='0': test_error",
		annotate(fmt.Errorf("test_error"), "TestAnnotate", "hello", "world", "id", 0).Error(),
	)
}

func TestAnnotateOdd(t *testing.T) {
	assert.Equal(t,
		"performing TestAnnotate hello='world' id='0' odd='<missing value>': test_error",
		annotate(fmt.Errorf("test_error"), "TestAnnotate", "hello", "world", "id", 0, "odd").Error(),
	)
}

func TestAnnotateNoKV(t *testing.T) {
	assert.Equal(t,
		"performing TestAnnotate: test_error",
		annotate(fmt.Errorf("test_error"), "TestAnnotate").Error(),
	)
}
