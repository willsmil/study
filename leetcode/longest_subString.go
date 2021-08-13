package leetcode

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	m := make(map[byte]int)
	l := 0
	r := 1
	max := 1
	m[s[0]] = 0
	for l <= r && r < len(s) {
		if _, ok := m[s[r]]; ok {
			delete(m, s[l])
			l++
		} else {
			m[s[r]] = r
			r++
			if r-l > max {
				max = r - l
			}
		}
	}
	return max
}
