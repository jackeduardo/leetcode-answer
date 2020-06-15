package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
}

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid //为什么不加1，因为right位置可能就是最大值
		} else {
			left = mid + 1 //为什么加一，因为left位置不可能是最大值
		}
	}
	return left
}
