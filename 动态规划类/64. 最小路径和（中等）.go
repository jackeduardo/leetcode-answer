package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}

//没有太大难度的dp路径题，建立一个m+1*n+1的dp二维数组，对应存储每个grid[i，j]的路径和，每次取最小值，那么dp[m-1][n-1]就是最小路径和
//dp[i][j]表示:到grid[i-1][j-1]的路径和
func minPathSum(grid [][]int) int {
	m := len(grid) + 1
	n := len(grid[0]) + 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[1][1] = grid[0][0]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if i == 1 {
				dp[i][j] = dp[i][j-1] + grid[i-1][j-1]
			} else if j == 1 {
				dp[i][j] = dp[i-1][j] + grid[i-1][j-1]
			} else {
				dp[i][j] = Methods.Min(dp[i][j-1]+grid[i-1][j-1], dp[i-1][j]+grid[i-1][j-1])
			}
		}
	}
	return dp[m-1][n-1]

}
