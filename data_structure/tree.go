package data_structure

import (
	"fmt"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func NewBSTTree(l []int) *TreeNode {
	if len(l) == 0 {
		return &TreeNode{}
	}
	root := NewTreeNode(l[0])
	for i := 1; i < len(l); i++ {
		root.BSTAdd(l[i])
	}
	return root
}

func NewTreeNode(i int) *TreeNode {
	return &TreeNode{data: i}
}

func (t *TreeNode) BSTAdd(i int) {
	n := TreeNode{data: i}
	var parent *TreeNode
	cur := t
	for cur != nil {
		parent = cur
		if i > cur.data {
			cur = cur.right
			if cur == nil {
				parent.right = &n
			}
		} else {
			cur = cur.left
			if cur == nil {
				parent.left = &n
			}
		}
	}
}

func (t *TreeNode) PrintLeft() {
	if t == nil {
		return
	}
	fmt.Println(t.data)
	t.left.PrintLeft()
	t.right.PrintLeft()
}

func (t *TreeNode) PrintRight() {
	if t == nil {
		return
	}
	t.left.PrintRight()
	t.right.PrintRight()
	fmt.Println(t.data)
}

func (t *TreeNode) PrintMid() {
	if t == nil {
		return
	}
	t.left.PrintMid()
	fmt.Println(t.data)
	t.right.PrintMid()
}

func (t *TreeNode) KthSmall(k int) {

}
