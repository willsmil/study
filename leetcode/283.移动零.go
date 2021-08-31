package leetcode

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[l] == 0 {
			if nums[r] != 0 {
				nums[l], nums[r] = nums[r], nums[l]
				l++
			}
		} else {
			l++
		}
	}
}

// @lc code=end
