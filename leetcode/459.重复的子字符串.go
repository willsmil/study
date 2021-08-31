package leetcode

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for l := 1; l < len(s)/2+1; l++ {
		if len(s)%l != 0 {
			continue
		}
		if IsCombine(s, s[:l]) {
			return true
		}
	}
	return false
}

func IsCombine(s, subStr string) bool {
	l := len(subStr)
	if len(s)%l != 0 {
		return false
	}
	for i := 0; i+l <= len(s); i += l {
		if s[i:i+l] != subStr {
			return false
		}
	}
	return true
}

// @lc code=end
