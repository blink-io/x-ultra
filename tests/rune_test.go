package tests

import (
	"fmt"
	"testing"
)

func TestRune_1(t *testing.T) {
	raw := "你好,Hello."
	var rr = []rune(raw)
	var bb = []byte(raw)

	fmt.Println("len of rune: ", len(rr))
	fmt.Println("len of utf8: ", len(bb))

	for _, v := range rr {
		fmt.Printf("%X\n", v)
	}
	//fmt.Println("rune slice: ", rr, ", len: ", len(rr))
}
