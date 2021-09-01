package leetcode

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回滑动窗口中的最大值。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//
//
// 示例 4：
//
//
//输入：nums = [9,11], k = 2
//输出：[11]
//
//
// 示例 5：
//
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 1151 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewQueue()
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	var res []int
	res = append(res, queue.Front())
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.Front())
	}
	return res
}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{data: []int{}}
}
func (q *Queue) Pop(x int) {
	if len(q.data) > 0 && q.data[0] == x {
		q.data = q.data[1:]
	}
}
func (q *Queue) Push(x int) {
	for len(q.data) > 0 && x > q.data[len(q.data)-1] {
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, x)
}
func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return 0
	}
	return q.data[0]
}

//leetcode submit region end(Prohibit modification and deletion)
