package main

import "fmt"

func main() {
	n:=5
	fmt.Println(fib2(n))
}
//超时递归
func fib(n int) int {
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	return fib(n-1)+fib(n-2)
}
//基础dp
func fib2(n int)  int{
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	dp:=make([]int,n+1)
	dp[0]=0
	dp[1]=1
	for i := 2; i<n+1; i++ {
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}
