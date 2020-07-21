package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"math"
)

func main() {
	fmt.Println(getMoneyAmount(20))
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
	}
	for i := n; i >= 1; i-- {
		for j := i; j <= n; j++ {
			if i == j {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt32
				for k := i; k <= j; k++ {
					dp[i][j] = Methods.Min(dp[i][j], Methods.Max(dp[i][k-1], dp[k+1][j])+k)
					fmt.Println(i, " ", j, " ", dp[i][j])
				}
			}
		}
	}
	return dp[1][n]
}
