package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

//一维dp,有一个坑点就是如果说选择间隔为2情况，那这个算法会不会存在遗漏这种情况更大的数值，例如取num[i],num[i+3]
//答案是不会的，因为如果数值是更大的那一种情况，它被会反映在max函数中，被继承到dp[i]中
//其实流程中就已经包含了num[i]+num[3]的可能性，如果它是更大的，他的状态会被继承的
func rob(nums []int) int {
	dp := make([]int, len(nums)+2)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < len(dp); i++ {
		dp[i] = Methods.Max(dp[i-1], nums[i-2]+dp[i-2])
	}
	return dp[len(dp)-1]
}
