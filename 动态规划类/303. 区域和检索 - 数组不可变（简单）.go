package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange_1D(0, 5))

}

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		arr: nums,
	}
}

//二维dp超内存了，不过应该是对的
func (this *NumArray) SumRange(i int, j int) int {
	length := len(this.arr)
	dp := make([][]int, length)
	for a := 0; a < length; a++ {
		dp[a] = make([]int, length)
	}
	dp[0][0] = this.arr[0]
	for b := 1; b < length; b++ {
		dp[0][b] = dp[0][b-1] + this.arr[b]
	}
	for a := 1; a < length; a++ {
		for b := 1; b < length; b++ {
			if a <= b {
				dp[a][b] = dp[a-1][b] - dp[a-1][a-1]
			}
		}
	}
	return dp[i][j]
}

//一维dp缓存
func (this *NumArray) SumRange_1D(i int, j int) int {
	length := len(this.arr)
	dp := make([]int, length+1)
	dp[0] = 0
	for a := 1; a <= length; a++ {
		dp[a] = dp[a-1] + this.arr[a-1]
	}
	return dp[j+1] - dp[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
