package leetcode

// 面试题 08.04. 幂集
// 幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
// 说明：解集不能包含重复的子集。
func subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(path []int, index int)

	dfs = func(path []int, index int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(path, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
