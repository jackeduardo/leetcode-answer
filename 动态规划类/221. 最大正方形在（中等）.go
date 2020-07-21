package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}
//如果定义dp[i][j]为i，j位置的最大正方形面积会非常不方便，因为很难处理边长这个变量，所以dp[i][j]定义为最大边长
func maximalSquare(matrix [][]byte) int {
	if len(matrix)==0{
		return 0
	}
	m:=len(matrix)
	n:=len(matrix[0])
	dp:=make([][]int,m+1)
	for i:=range dp{
		dp[i]=make([]int,n+1)
	}
	side:=0
	for i :=1; i <=m; i++ {
		for j := 1; j <=n; j++ {
			if matrix[i-1][j-1]=='1'{
				dp[i][j]=1+Methods.Min(dp[i-1][j-1],Methods.Min(dp[i][j-1],dp[i-1][j]))//因为是正方形，所以要取min
				side=Methods.Max(side,dp[i][j])
			}
		}
	}
	return side*side
}
