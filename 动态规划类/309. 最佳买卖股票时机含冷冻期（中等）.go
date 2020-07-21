package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit2(prices))
}

//不管是买还是卖，买卖合在一起是一次交易，所以只要买或者卖的行为发生一次，就是一次交易（买/卖取一种就可以），不能买卖都算一次交易，不然就变成2次交易了
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	pre := 0
	for i := 1; i < len(prices); i++ {
		temp := dp[i-1][0]
		dp[i][0] = Methods.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Methods.Max(dp[i-1][1], pre-prices[i])
		pre = temp
	}
	return dp[len(prices)-1][0]
}

//三个状态dp
func maxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	sell := make([]int, len(prices))
	buy := make([]int, len(prices))
	cooldown := make([]int, len(prices))
	buy[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		sell[i] = Methods.Max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = Methods.Max(buy[i-1], cooldown[i-1]-prices[i])
		cooldown[i] = Methods.Max(cooldown[i-1], sell[i-1])
	}
	return Methods.Max(sell[len(prices)-1], cooldown[len(prices)-1])
}

func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
