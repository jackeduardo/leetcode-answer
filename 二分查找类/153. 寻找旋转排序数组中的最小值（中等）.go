package main

import (
	"fmt"
)

func main() {
	nums := []int{4,5,1,2,3}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid//mid位置可能是最小值，所以保留
		} else {
			left = mid + 1//mid位置肯定不是最小值，所以mid+1
		}
	}
	return nums[left]
}
