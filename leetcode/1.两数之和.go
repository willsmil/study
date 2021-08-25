package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 1. 暴力解法
	// if len(nums) == 0 {
	// 	return nil
	// }
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return nil

	// 2. 哈希表
	m := make(map[int]int)
	for i, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{v, i}
		} else {
			m[n] = i
		}
	}
	return nil
}

// @lc code=end
