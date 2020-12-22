package leetcode

import (
	"fmt"
	"testing"
)

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
