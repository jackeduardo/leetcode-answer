package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

//二分法

func searchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)
	for left != right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
