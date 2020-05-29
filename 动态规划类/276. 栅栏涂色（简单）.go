package main

import "fmt"

func main() {
	fmt.Println(numWays(3, 2))
}

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}
	dp := make([]int, n)
	dp[0] = k
	dp[1] = k * k
	for i := 2; i <= n-1; i++ {
		dp[i] = dp[i-2]*(k-1) + dp[i-1]*(k-1)
	}
	return dp[n-1]
}
