package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	prices := []int{2, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfitDp(prices))
}

//贪心算法，只要今天价格高于昨天就卖，获取最大利润，依据可以通过 a<b<c<d推导
func maxProfit(prices []int) int {
	n := len(prices)
	profit := 0
	for l, r := 0, 1; r < n; {
		if prices[r] > prices[l] {
			profit += prices[r] - prices[l]
		}
		l++
		r++
	}
	return profit
}

//dp算法 比较搞，不推荐。 dp方法主要就是记录所有的状态，但是这种找所有状态的记录实在太难找了
func maxProfitDp(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = Methods.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Methods.Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
