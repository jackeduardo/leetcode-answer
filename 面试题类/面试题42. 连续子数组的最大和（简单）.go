package main

import (
	"fmt"
	"leetcode-answer/Methods"
)

func main() {
	nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray2(nums))
}
//一维dp正解,优化
func maxSubArray(nums []int) int {
	if nums==nil{
		return 0
	}
	dp:=make([]int,len(nums))
	dp[0]=nums[0]
	max:=nums[0]
	for i := 1; i<len(nums); i++ {
		dp[i]=Methods.Max(nums[i],dp[i-1]+nums[i])
		max=Methods.Max(max,dp[i])
	}
	return max
}
//一维超时，是下面二维的优化版省了内存，但是还是全遍历爆算了，依然超时
func maxSubArray2(nums []int) int {
	if nums==nil{
		return 0
	}
	n:=len(nums)
	max:=nums[0]
	dp:=make([]int,n)
	//dp[0]=nums[0]
	for i := 0; i<n; i++ {
		for j:=0;j<=i;j++{
			if i==j{
				dp[j]=nums[i]
			}else{
				dp[j]=dp[j]+nums[i]
			}
			max=Methods.Max(max,dp[j])
		}
	}
	return max
}
//二维dp超内存
func maxSubArray_2D(nums []int) int {
	if nums==nil{
		return 0
	}
	n:=len(nums)
	max:=nums[0]
	dp:=make([][]int,n)
	for i := 0; i <n; i++ {
		dp[i]=make([]int,n)
	}
	dp[0][0]=nums[0]
	for i := 1; i<n; i++ {
		for j:=0;j<=i;j++{
			if i==j{
				dp[i][j]=nums[i]
			}else{
				dp[i][j]=dp[i-1][j]+nums[i]
			}
			max=Methods.Max(max,dp[i][j])
		}
	}
	return max

}
