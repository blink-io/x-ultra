package http3

import (
	"fmt"
	"testing"
)

func TestKind(t *testing.T) {
	var a *adapter
	fmt.Println("Kind: ", a.Kind())
}
