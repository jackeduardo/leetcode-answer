package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}
//二维数组记录所有状态，dp[i][j]表示i到j位置能获得的最大数量
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 2 {
		return nums[0]
	}
	nums = append(append([]int{1}, nums...), 1)
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums))
	}
	for i := 2; i < len(nums); i++ {//i表示相对于j往后的位置步数
		for j := 0; j+i < len(nums); j++ {
			for k := j + 1; k < i+j; k++ {
				dp[j][j+i] = Methods.Max(dp[j][j+i], dp[j][k]+dp[k][j+i]+nums[j]*nums[k]*nums[j+i])
			}
		}
	}
	return dp[0][len(dp)-1]
}
