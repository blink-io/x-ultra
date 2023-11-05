package tests

import (
	"fmt"
	"testing"

	"github.com/sanity-io/litter"
)

func TestLitter_1(t *testing.T) {
	type Result struct {
		Action  string
		Level   int
		Score   float32
		Enabled bool
	}
	res := &Result{
		Action:  "moveon",
		Level:   111,
		Score:   222.222,
		Enabled: true,
	}
	actual := litter.Sdump(res)

	fmt.Println("result: ", actual)
}
