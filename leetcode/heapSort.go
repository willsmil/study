package leetcode

import "fmt"

// root = idx0
// left = root*2+1
// right = root*2+2
// parent = (cur-1)/2
func heapSort(nums []int) []int {
	x := nums[:1]
	for i := 1; i < len(nums); i++ {
		x = append(x, nums[i])
		_sort(x)
		fmt.Println(x)
	}
	return x
}

func _sort(nums []int) {
	cur := len(nums) - 1
	parent := (cur - 1) / 2
	for cur != 0 && nums[cur] > nums[parent] {
		nums[cur], nums[parent] = nums[parent], nums[cur]
		cur = parent
		parent = (cur - 1) / 2
	}
}
