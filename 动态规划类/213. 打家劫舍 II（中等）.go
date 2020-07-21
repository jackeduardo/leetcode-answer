package main

import (
	"fmt"
)

func main() {
	nums:=[]int{2,3,2}
	fmt.Println(robII(nums))
}
//第一次0到n-2，第二次1到n-1，两次dp，取最大值
func robII(nums []int) int {
	if len(nums)==0{
		return  0
	}
	if len(nums)==1{
		return nums[0]
	}
	dp1 := make([]int, len(nums)+1)
	for i := 2; i < len(nums)+1; i++ {
		dp1[i] = Max(dp1[i-1], nums[i-2]+dp1[i-2])
	}
	dp2 := make([]int, len(nums)+1)
	for i := 2; i < len(nums)+1; i++ {
		dp2[i] = Max(dp2[i-1], nums[i-1]+dp2[i-2])
	}
	return Max(dp1[len(dp1)-1],dp2[len(dp2)-1])
}