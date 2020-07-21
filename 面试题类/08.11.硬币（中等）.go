package main

import "fmt"

//动态规划
//回溯感觉也可以做，但是时间复杂度太高，因为要遍历所有情况，会有很多重复情况
func main() {
	n := 10
	fmt.Print(waysToChange1D(n))
}

//二维（基础）
func waysToChange(n int) int {
	coins := []int{1, 5, 10, 25}
	dp := make([][]int, 4)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	//默认值为1，因为任何n都至少有一种表达结果（全是1），n>=1
	for i := range dp[0] {
		dp[0][i] = 1
	}
	//为什么外层循环用coins，为了去重10+5和5+10是一种组合的情况
	for i := 1; i < 4; i++ {
		for j := 1; j < n+1; j++ {
			if j >= coins[i] {
				dp[i][j] = (dp[i-1][j] + dp[i][j-coins[i]]) % 1000000007 //核心递推公式，表示：如果不取当前硬币就继承上个硬币对应的方法集结果，如果取当前硬币那么方法集为j分减去当前硬币（意思是取到还剩当前一次硬币数的所有方法）
			} else { //这其实是一个数学的化简结果，用数学推到更好理解。如果直接从公式上去解释并不好理解，只能抽象的理解为取或者不取的方法集，层层嵌套
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[3][n]
}

//一维
//利用初始化全为1的dp[]一维数组，将二维dp化简为一维dp
func waysToChange1D(n int) int {
	coins := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < 4; i++ {
		for j := 1; j < n+1; j++ {
			if j >= coins[i] {
				dp[j] = (dp[j] + dp[j-coins[i]]) % 1000000007 //要是来单单理解这段公式挺困难的，其实就是上面二维dp的化简，硬币的维度隐藏在了重复遍历一个数组的过程中，本质还是上面的逻辑，因为对于一维数组来说，旧数据
			} //dp[j]就是二维dp中的dp[i-1][j]

		}
	}
	return dp[n]
}
