package tests

import (
	"fmt"
	"testing"

	"github.com/sourcegraph/conc"
)

func TestConc_1(t *testing.T) {
	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			fmt.Println("Invoke in Go func")
		})
	}
	wg.Wait()
}
