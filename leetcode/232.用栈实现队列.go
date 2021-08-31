package leetcode

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	cur *Stack // 存正序
	tmp *Stack // 存逆序
	// true:cur有效，false:tmp有效
	flag bool
}

/** Initialize your data structure here. */
func NewQuene() MyQueue {
	return MyQueue{cur: &Stack{}, tmp: &Stack{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.flag {
		this.cur.Push(x)
		return
	}
	for !this.tmp.IsEmpty() {
		this.cur.Push(this.tmp.Pop())
	}
	this.cur.Push(x)
	this.flag = true
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.flag {
		return this.tmp.Pop()
	}
	for !this.cur.IsEmpty() {
		this.tmp.Push(this.cur.Pop())
	}
	this.flag = false
	return this.tmp.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.flag {
		return this.tmp.Peek()
	}
	for !this.cur.IsEmpty() {
		this.tmp.Push(this.cur.Pop())
	}
	this.flag = false
	return this.tmp.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.flag {
		return this.cur.IsEmpty()
	}
	return this.tmp.IsEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type Stack struct {
	data []int
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// @lc code=end
