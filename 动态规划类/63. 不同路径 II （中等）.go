package main

import "fmt"

func main() {
	obstacleGrid:=[][]int{
		{1,1},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
//写的繁琐了点，但是非常好理解，先把所有1变成-1
//初始化第一行第一列的时候，把-1前面的数全变成1，后面的数保持为0
//最后把第一行第一列中的-1全变成0即可，完成初始化工作
//之后按照dp[i][j]=dp[i-1][j]+dp[i][j-1]完成求解
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m:=len(obstacleGrid)
	n:=len(obstacleGrid[0])
	for i := 0; i <m; i++ {
		for j := 0; j <n; j++ {
			if obstacleGrid[i][j]==1{
				obstacleGrid[i][j]=-1
			}
		}
	}
	for i := 0; i <m; i++ {
		if obstacleGrid[i][0]==-1{
			break
		}
		obstacleGrid[i][0]=1
	}
	for j := 0; j <n; j++ {
		if obstacleGrid[0][j]==-1{
			break
		}
		obstacleGrid[0][j]=1
	}
	for i := 0; i <m; i++ {
		if obstacleGrid[i][0]==-1{
			obstacleGrid[i][0]=0
		}
	}
	for j := 0; j <n; j++ {
		if obstacleGrid[0][j]==-1{
			obstacleGrid[0][j]=0
		}
	}

	for i := 1; i<m; i++ {
		for j := 1; j<n; j++ {
			if obstacleGrid[i][j]==-1{
				obstacleGrid[i][j]=0
				continue
			}
			obstacleGrid[i][j]=obstacleGrid[i-1][j]+obstacleGrid[i][j-1]
		}
	}
	return obstacleGrid[m-1][n-1]
}