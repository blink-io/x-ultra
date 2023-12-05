package attrs

import (
	"fmt"
	"testing"
)

func TestAttrs_1(t *testing.T) {
	attrs := Make("one", 1, "one", 11)
	fmt.Println(attrs)
}
