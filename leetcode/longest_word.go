package leetcode

import (
	"fmt"
	"strings"
)

// 面试题 17.15. 最长单词
// 给定一组单词words，编写一个程序，找出其中的最长单词，且该单词由这组单词中的其他单词组合而成。
// 若有多个长度相同的结果，返回其中字典序最小的一项，若没有符合要求的单词则返回空字符串。
//
// 输入： ["cat","banana","dog","nana","walk","walker","dogwalker"]
// 输出： "dogwalker"
// 解释： "dogwalker"可由"dog"和"walker"组成。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-word-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func longestWord(words []string) string {
	length := 0
	key := ""
	var contains []string
	for i, word := range words {
		contains = []string{}
		for j := range words {
			if i == j {
				continue
			}
			if strings.Contains(word, words[j]) {
				contains = append(contains, words[j])
			}
		}
		if canBeCombined(word, contains) {
			if len(word) > length {
				key = word
				length = len(word)
			} else if len(word) == length {
				if word < key {
					key = word
				}
			}
		}
	}
	return key
}
func canBeCombined(word string, words []string) bool {
	tmp := word
	for range words {
		if trimStart(tmp, words) {
			return true
		}
	}
	return false
}
func trimStart(s string, words []string) bool {
	fmt.Println("trim", s, words)
	if s == "" {
		return true
	}
	for _, w := range words {
		if strings.Index(s, w) == 0 {
			if trimStart(s[len(w):], words) {
				return true
			}
		}
	}
	return false
}
