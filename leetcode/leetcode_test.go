package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
}

func TestSubsets(t *testing.T) {
	input := []int{1, 2, 3}
	fmt.Println(subsets(input))
}

func TestFindRestaurant(t *testing.T) {
	l1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	l2 := []string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant(l1, l2))
}

func TestLongestWords(t *testing.T) {
	words := []string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"}
	fmt.Println(longestWord(words))
}

func TestSequentialDigits(t *testing.T) {
	fmt.Println(sequentialDigits(10, 200))
}

func TestAddTwoNums(t *testing.T) {
	var h1, h2 *ListNode
	h1 = &ListNode{Val: 0, Next: nil}
	h2 = &ListNode{Val: 9, Next: nil}
	// h1.Next = &ListNode{Val: 4, Next: nil}
	// h1.Next.Next = &ListNode{Val: 4, Next: nil}
	h2.Next = &ListNode{Val: 1, Next: nil}
	h2.Next.Next = &ListNode{Val: 6, Next: nil}
	res := addTwoNumbers(h1, h2)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Println("nil")
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
