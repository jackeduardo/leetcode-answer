package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"math"
)

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

//动态规划，由于正负性质的原因，必须同时每个状态下最大值和最小值
func maxProduct(nums []int) int {
	res := math.MinInt32
	min, max := 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}
		min = Methods.Min(min*nums[i], nums[i])
		max = Methods.Max(max*nums[i], nums[i])
		res = Methods.Max(res, max)
	}
	return res
}
