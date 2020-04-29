package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Print(canJump1(nums))
}

//暴力贪心，遍历所有位置，维护最长可达位置
func canJump1(nums []int) bool {
	rightmost := 0
	for i := 0; i < len(nums); i++ {
		if i > rightmost {
			return false
		}
		rightmost = Methods.Max(rightmost, i+nums[i])
	}
	return true
}
