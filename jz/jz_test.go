package jz

import (
	"fmt"
	"testing"
)

func TestLeftRotateString(t *testing.T) {
	x := LeftRotateString("abcde", 2)
	t.Log(x)
}

func TestSliceInsert(t *testing.T) {
	s := []int{1, 3, 4}
	ss := append(s[:2], 2)
	fmt.Println(s)
	s = append(ss, s[2:]...)
	fmt.Println(s)
}
