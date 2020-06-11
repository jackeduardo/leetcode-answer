package main

import "fmt"

func main() {
	nums:=[]int{1,2,2,3}
	fmt.Println(removeDuplicates(nums))
}
 //双指针去重，要点是设定一个index依赖于for循环中的i，只有当前两位的时候或者当前nums[i]大于nums[index-2],因为观察符合条件的数组就能发现，对于任意i>=2的情况下，nums[i]必须大于nums[i-2]
func removeDuplicates(nums []int) int {
	index:=0
	for i := 0; i <len(nums); i++ {
		if index<2||nums[i]>nums[index-2]{
			nums[index]=nums[i]
			index++
		}
	}
	return index
}
