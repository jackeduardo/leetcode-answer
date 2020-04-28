package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Print(maxSubArray(nums))
}

//维护一个max取sum的最大值，坑点在于题目要求连续的子序合，所以sum为负的时候就直接放弃，取新值，而对于max来讲是没有影响的，哪怕nums[i]的值更小，max里面储存的也是最大的值
//利用动态规划来做，是一个思路
func maxSubArray(nums []int) int {
	max:=nums[0]
	sum:=nums[0]
	for i := 1; i <len(nums); i++ {
		if sum<0{
			sum=nums[i]
		}else{
			sum+=nums[i]
		}
		max=Methods.Max(max,sum)
	}
	return max
}