package tests

import (
	"fmt"
	"testing"
)

func TestSlice_1(t *testing.T) {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	handleSlice(s1)
	fmt.Println(s1)
}

func TestSlice_2(t *testing.T) {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	handleSlice2(&s1)
	fmt.Println(s1)
}

func handleSlice(s []int) {
	s[2] = 22
	s[1] = 11
}

func handleSlice2(s *[]int) {
	(*s)[2] = 44
	(*s)[1] = 55
}
