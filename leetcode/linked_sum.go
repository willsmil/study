package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	add := 0
	val := 0
	cur := head
	var single *ListNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			single = l2
			break
		} else if l2 == nil {
			single = l1
			break
		}
		val = l1.Val + l2.Val + add
		if val >= 10 {
			cur.Next = &ListNode{Val: val - 10, Next: nil}
			add = 1
		} else {
			cur.Next = &ListNode{Val: val, Next: nil}
			add = 0
		}
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}
	for ; single != nil; single = single.Next {
		val = single.Val + add
		if val < 10 {
			single.Val = val
			cur.Next = single
			add = 0
			break
		}
		cur.Next = &ListNode{Val: 0, Next: nil}
		cur = cur.Next
		add = 1
	}
	if add > 0 {
		cur.Next = &ListNode{Val: 1, Next: nil}
	}
	return head.Next
}
