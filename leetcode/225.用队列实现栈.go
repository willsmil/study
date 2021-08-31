package leetcode

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {

}

/** Get the top element. */
func (this *MyStack) Top() int {

}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {

}

type Quene struct {
	data []int
}

func (q *Quene) Push(x int) {
	q.data = append(q.data, x)
}

func (q *Quene) Pop() int {
	if q.Empty() {
		return 0
	}
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Quene) Top() int {
	if q.Empty() {
		return 0
	}
	return q.data[0]
}

func (q *Quene) Empty() bool {
	return len(q.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
