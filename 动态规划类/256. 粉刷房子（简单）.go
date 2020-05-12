package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	costs := [][]int{
		{5, 8, 6},
		{19, 14, 13},
		{7, 5, 12},
		{14, 15, 17},
		{3, 20, 10},
	}
	fmt.Println(minCost(costs))
}

//一开始写的算法dp是错的，写的是贪心算法，只考察了局部最优，但不是全局最优，所以依然熟悉dp的核心，dp是全局最优
//这里的 costs数组其实就是dp数据，只是重复创建一个dp了，代表的是cost[i][j]的最小成本，全局的最小成本是最后一个单位中的最小元素，这个二维数组中每一个0/1/2位置都是排除了上一个index里的相同位置的成本
func minCost(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	for i := 1; i < len(costs); i++ {
		costs[i][0] += Methods.Min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += Methods.Min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += Methods.Min(costs[i-1][0], costs[i-1][1])
	}
	return Methods.Min(Methods.Min(costs[n-1][0], costs[n-1][1]), costs[n-1][2])
}
