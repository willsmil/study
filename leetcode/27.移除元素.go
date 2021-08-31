package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 && nums[0] == val {
		return 0
	}
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			n++
		} else {
			if n != 0 {
				nums[i], nums[i-n] = nums[i-n], nums[i]
			}
		}
	}
	return len(nums) - n
}

// @lc code=end
