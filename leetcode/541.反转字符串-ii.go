package leetcode

/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	arr := []byte(s)
	l := len(arr)
	for i := 0; i < l; i = i + 2*k {
		if l-i < k {
			reverse(arr, i, l-1)
		} else {
			reverse(arr, i, i+k-1)
		}
	}
	return string(arr)
}

func reverse(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// @lc code=end
