/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

package leetcode

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 || len(nums3) == 0 || len(nums4) == 0 {
		return 0
	}
	var res int
	m := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if v, ok := m[0-n3-n4]; ok {
				res += v
			}
		}
	}
	return res
}

// @lc code=end
