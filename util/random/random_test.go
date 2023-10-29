package random

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Len(t, String(32), 32)
	r := New()
	str := r.String(8, Numeric)
	fmt.Println("rand str:  ", str)
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), str)
}
