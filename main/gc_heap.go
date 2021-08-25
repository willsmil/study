package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	// x := ret()
	var s []int
	s = append(s, 3)
	s = append(s, 2)
	s = append(s, 1)
	s = append(s, 4)
	s = append(s, 5)
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)
}

func ret() *int {
	a := 1
	p := unsafe.Pointer(&a)
	return (*int)(noescape(p))
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x^0)
}