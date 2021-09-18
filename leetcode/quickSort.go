package leetcode

import "fmt"

func quickSort(nums []int) []int {
	if len(nums) > 1 {
		x := find(nums)
		fmt.Println(x, nums)
		if x > 1 {
			quickSort(nums[:x])
		}
		if x < len(nums)-2 {
			quickSort(nums[x+1:])
		}
	}
	return nums
}

func find(nums []int) int {
	l, r := 1, len(nums)-1
	c := nums[0]
	for l != r {
		for l < r && nums[l] > c {
			l++
		}
		for l < r && nums[r] < c {
			r--
		}
		if l < r {
			nums[r], nums[l] = nums[l], nums[r]
		}
	}
	if nums[l] > c {
		nums[0], nums[l] = nums[l], nums[0]
		return l
	}
	nums[0], nums[l-1] = nums[l-1], nums[0]
	return l - 1
}
