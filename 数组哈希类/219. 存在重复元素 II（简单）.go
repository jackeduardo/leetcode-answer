package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}

// 哈希表取重,哈希表每次都要更新以免漏掉
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			if Methods.Abs(i-m[nums[i]]) <= k {
				return true
			}
			m[nums[i]] = i
		} else {
			m[nums[i]] = i
		}
	}
	return false
}
