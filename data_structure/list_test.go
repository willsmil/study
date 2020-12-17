package data_structure

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	var head SinglyLinkedNode
	for i := 0; i < 10; i++ {
		head.Add(i)
	}
	x := head.ReverseWithHead()
	x.Print()
	fmt.Println()
	head.Print()

}
