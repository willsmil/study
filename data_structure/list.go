package data_structure

import "fmt"

// 单项链表
// 如何实现一个高效的单向链表逆序输出？
type SinglyLinkedNode struct {
	Data int
	Next *SinglyLinkedNode
}

func New() *SinglyLinkedNode {
	return &SinglyLinkedNode{}
}

func (n *SinglyLinkedNode) Add(i int) {
	t := n
	for t.Next != nil {
		t = t.Next
	}
	t.Next = &SinglyLinkedNode{Data: i}
}

func (n *SinglyLinkedNode) AddHead(i int) {
	t := &SinglyLinkedNode{Data: i}
	t.Next = n.Next
	n.Next = t
}

func (n *SinglyLinkedNode) Print() {
	t := n
	for t != nil {
		fmt.Print(t.Data, "->")
		t = t.Next
	}
}

func (n *SinglyLinkedNode) ReverseWithSpace() {
	var ns []int
	t := n
	for t != nil {
		ns = append(ns, t.Data)
		t = t.Next
	}
	for i := len(ns) - 1; i > -1; i-- {
		fmt.Print(ns[i], "->")
	}
}
func (n *SinglyLinkedNode) ReverseWithDelete() {
	head := &SinglyLinkedNode{}
	head.Next = n.Next
	t := head
	for t.Next != nil {
		if t.Next.Next == nil {
			fmt.Print(t.Next.Data, "->")
			t.Next = nil
			break
		}
		for t.Next.Next != nil {
			t = t.Next
		}
		fmt.Print(t.Next.Data, "->")
		t.Next = nil
		t = n
	}
	fmt.Print(t.Data, "->")
}
func (n *SinglyLinkedNode) ReverseWithHead() *SinglyLinkedNode {
	var head SinglyLinkedNode
	t := n
	for t.Next != nil {
		head.AddHead(t.Data)
		t = t.Next
	}
	head.Data = t.Data
	return &head
}
func (n *SinglyLinkedNode) Reverse() {
	if n == nil {
		return
	}

	var pLeft *SinglyLinkedNode
	current := n
	pRight := n.Next

	for pRight != nil {
		current.Next = pLeft
		tmp := pRight.Next
		pRight.Next = current
		pLeft = current
		current = pRight
		pRight = tmp
	}
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
}
