package leetcode

import "fmt"
/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var res [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp := []int{nums[i], nums[j], nums[k]}
					fmt.Println(tmp)
					if !exist(res, tmp) {
						res = append(res, tmp)
					}
				}
			}
		}
	}
	return res
}

func exist(arr [][]int, elem []int) bool {
	for _, a := range arr {
		if sliceEqual(a, elem) {
			return true
		}
	}
	return false
}

func sliceEqual(s1, s2 []int) bool {
	m1 := make(map[int]int, len(s1))
	for _, s := range s1 {
		m1[s]++
	}
	for _, s := range s2 {
		m1[s]--
		if m1[s] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end
