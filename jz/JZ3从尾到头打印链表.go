package jz

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	var res []int
	for ; head != nil; head = head.Next {
		res = append([]int{head.Val}, res...)
	}
	return res
}
