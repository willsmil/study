package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	// 1. 暴力
	// ss := strings.Split(s, " ")
	// if len(ss) == 1 {
	// 	return s
	// }
	// var res string
	// for i := len(ss) - 1; i >= 0; i-- {
	// 	if ss[i] == "" {
	// 		continue
	// 	}
	// 	res += " " + ss[i]
	// }

	// return res[1:]
	// 2.---
	s = strings.TrimSpace(s)
	ss := []byte(s)
	l := len(ss)
	for i := 0; i < l/2; i++ {
		ss[i], ss[l-i] = ss[l-i], ss[i]
	}
	for i := 0; i < l; i++ {
		j := i
		for j := i + 1; ss[j] != ' ' || j == l; j++ {
		}
		if j-i == 1 {
			ss = append(ss[:i], ss[i+1:]...)
		} else {
			reverse1(ss, i, j-1)
			i = j
		}
	}
	return string(ss)
}

func reverse1(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// @lc code=end
