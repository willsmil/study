package leetcode

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, s := range magazine {
		m[s]++
	}
	for _, s := range ransomNote {
		if v, ok := m[s]; !ok || v == 0 {
			return false
		} else {
			m[s]--
		}
	}
	return true
}

// @lc code=end
