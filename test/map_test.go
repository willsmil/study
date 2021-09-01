package test

import (
	"fmt"
	"testing"
)

func TestMapRange(t *testing.T) {
	m := make(map[int]struct{}, 10)
	for i := 0; i < 10; i++ {
		m[i] = struct{}{}
	}
	for i := 0; i < 10; i++ {
		for k := range m {
			fmt.Print(k, ",")
		}
		fmt.Println()
	}
}
