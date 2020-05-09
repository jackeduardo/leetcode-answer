package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

//dp解法

func climbStairs(n int) int {
	dp:=make([]int,n+1)
	dp[0]=1
	dp[1]=1
	for i := 2; i<=n; i++ {
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}

//递归解法，超时
func climbStairs_rec(n int) int {
	if n==0{
		return 1
	}
	if n==1{
		return 1
	}
	return climbStairs_rec(n-1)+climbStairs_rec(n-2)
}