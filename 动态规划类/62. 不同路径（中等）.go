package main

import "fmt"

func main() {
	fmt.Println(uniquePaths_1D(5,4))
}
//二维dp，易于理解，dp[i][j]存储i，j位置的路径总和
//dp的第一行和第一列都只有一条路径，所以初始化为1
//递推公式 dp[i][j]=dp[i-1][j]+dp[i][j-1]
func uniquePaths(m int, n int) int {
	dp:=make([][]int,m)
	for i, _ := range dp {
		dp[i]=make([]int,n)
	}
	for i := 0; i <m; i++ {
		dp[i][0]=1
	}
	for j := 0; j <n; j++ {
		dp[0][j]=1
	}
	for i := 1; i <m; i++ {
		for j := 1; j <n; j++ {
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//优化过后的一维dp，因为每一行的信息只和上一行 /上一列有关
//dp[i]=dp[i-1]+dp[i]
//抛弃了不需要储存的信息
//相对难理解，建议上一种解法
func uniquePaths_1D(m int, n int) int {
	dp:=make([]int,n)
	for j := 0; j <n; j++ {
		dp[j]=1
	}
	for i := 1; i <m; i++ {
		dp[0]=1
		for j := 1; j <n; j++ {
			dp[j]=dp[j-1]+dp[j]
		}
	}
	return dp[n-1]
}