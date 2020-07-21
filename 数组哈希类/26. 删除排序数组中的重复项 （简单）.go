package main

import "fmt"

func main() {
	nums:=[]int{1,1,2}
	removeDuplicates(nums)
	fmt.Println(nums)
}
//双指针遍历
func removeDuplicates(nums []int) int {
	if len(nums)<2{
		return len(nums)
	}
	i,j:=0,1
	for j<len(nums){
		if nums[i]==nums[j]{
			j++
		}else {
			i++
			nums[i]=nums[j]
			j++
		}
	}
	return i+1
}
