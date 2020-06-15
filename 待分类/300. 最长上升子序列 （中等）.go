package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	nums:=[]int{4,10,4,3,8,9}
	fmt.Println(lengthOfLIS(nums))
}
//状态定义一定要清晰，这里的dp[i]是指在可构成上升序列的情况下的最大上升序列长度，如果不构成那当前位置的最大长度就为1
//等效理解为dp[i]的值代表以nums[i]结尾的最长子序列长度
func lengthOfLIS(nums []int) int {
	if len(nums)==0{
		return 0
	}
	dp:=make([]int,len(nums))
	for i := 0; i <len(dp); i++ {
		dp[i]=1
	}
	dp[0]=1
	max:=1
	for i := 1; i<len(dp); i++ {
		for j := 0; j <i; j++ {
			if nums[j]<nums[i]{
				dp[i]=dp[j]+1
			}
			max=Methods.Max(max,dp[i])
		}
	}
	return max
}