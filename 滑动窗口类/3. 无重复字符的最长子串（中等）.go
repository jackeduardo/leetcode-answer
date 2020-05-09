package main

import (
	"fmt"

	"leetcode-answer/Methods"
)

func main() {
	str := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(str))
}

//双指针法，滑动窗口思路配合哈希表
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left := 0
	right := 1
	set := make(map[byte]int)
	set[s[left]] = left
	maxlen := 1
	for ; right < len(s); right++ {
		cur := s[right]
		if _, ok := set[cur]; ok {
			left = Methods.Max(set[cur]+1, left) //这里为坑点，意思为如果重复值在left左边就放弃，取left，如果在右边那么就要更新为被重复位置加一
		}
		set[cur] = right
		maxlen = Methods.Max(maxlen, right-left+1)

	}
	return maxlen
}
