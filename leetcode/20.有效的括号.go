package leetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	var i, j int
	for i = 0; i < len(s)-1; {
		j = i + 1
		for ; j < len(s); j++ {
			if match(s[i], s[j]) {
				if j-i == 1 {
					break
				}
				if !isValid(s[i+1 : j]) {
					return false
				}
			}
		}
		if !match(s[i], s[j]) {
			return false
		}
		i = j + 1
	}
	if !match(s[i], s[j]) {
		return false
	}
	return true
}

func match(a, b byte) bool {
	switch a {
	case '(':
		if b == ')' {
			return true
		}
	case '[':
		if b == ']' {
			return true
		}
	case '{':
		if b == '}' {
			return true
		}
	}
	return false
}

// @lc code=end
