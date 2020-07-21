package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
}
func numSquares(n int) int {
	dp:=make([]int,n+1)
	for i := 0; i <len(dp); i++ {
		dp[i]=math.MaxInt32
	}
	dp[0]=0
	dp[1]=1
	for i := 1; i<n+1; i++ {
		for j := 1; j*j <=i; j++ {
			dp[i]=Methods.Min(dp[i],dp[i-j*j]+1)
		}
	}
	return dp[n]
}

