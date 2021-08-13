package leetcode

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	res := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			l := i
			r := j
			for s[l] == s[r] && l < r {
				l++
				r--
			}
			if l >= r && j+1-i > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
