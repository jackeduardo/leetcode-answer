package main

import (
	"fmt"
	"sort"
)
//切片只是对原数组的映射，尤其是append操作，要小心更改底层数组的值
func main() {

	costs := [][]int{
		{3,20,7,7,16,8,7,12,11,19,1},
			{10,14,3,3,9,13,4,12,14,13,1},
			{10,1,14,11,1,16,2,7,16,7,19},
			{13,20,17,15,3,13,8,10,7,8,9},
			{4,14,18,15,11,9,19,3,15,12,15},
			{14,12,16,19,2,12,13,3,11,10,9},
			{18,12,10,16,19,9,18,4,14,2,4},

	}
	fmt.Println(minCostII2(costs))
}
//暴力dp，速度比较慢
func minCostII(costs [][]int) int {
	n := len(costs)
	k:=len(costs[0])
	if n == 0 ||k==0{
		return 0
	}
	dp:=make([][]int,n)
	for i := 0; i <len(dp); i++ {
		dp[i]=make([]int,k)
	}
	copy(dp[0],costs[0])

	for i := 1; i<n; i++ {
		for j := 0; j <k; j++ {
			temp:=make([]int,len(dp[i-1]))
			copy(temp,dp[i-1])
			dp[i][j]=getMinCost(append(temp[:j],temp[j+1:]...))+costs[i][j]
		}
	}

	return getMinCost(dp[n-1])
}

func getMinCost(arr []int)int{
	temp:=make([]int,len(arr))
	copy(temp,arr)
	sort.Ints(temp)
	return temp[0]
}
//优化过后的O(kn)算法
func minCostII2(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	k:=len(costs[0])
	min1,min2:=-1,-1
	for i := 0; i <n; i++ {
		last1,last2:=min1,min2
		min1,min2=-1,-1
		for j := 0; j <k; j++ {
			if j!=last1{
				if last1>=0{
					costs[i][j]+=costs[i-1][last1]
				}
			}else {
				if last2>=0{
					costs[i][j]+=costs[i-1][last2]
				}
			}
			if min1<0||costs[i][j]<costs[i][min1]{
				min2,min1=min1,j//更新最小和第二小的值
			}else if min2<0||costs[i][j]<costs[i][min2]{
				min2=j
			}
		}
	}
	return costs[n-1][min1]
}