package leetcode

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	m := make(map[int]struct{})
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		} else {
			m[n] = struct{}{}
		}
		n = getNum(n)
	}
	return true
}

func getNum(n int) int {
	var res int
	for n > 0 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	return res
}

// @lc code=end
