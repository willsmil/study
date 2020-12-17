package data_structure

import (
	"fmt"
	"testing"
)

func TestTreeNode_BSTAdd(t *testing.T) {

	l := []int{10, 15, 9, 20, 35, 1, 6, 21}

	root := NewBSTTree(l)
	root.PrintLeft()
	fmt.Println("====")
	root.PrintRight()
	fmt.Println("====")
	root.PrintMid()
}
