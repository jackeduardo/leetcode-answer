package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		min = Methods.Min(prices[i], min)
		dp[i] = Methods.Max(dp[i-1], prices[i]-min)
	}
	return dp[len(prices)-1]

}
