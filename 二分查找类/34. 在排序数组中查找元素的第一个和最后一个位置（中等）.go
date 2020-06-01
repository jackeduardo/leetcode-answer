package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}

//代码写的稍微繁琐了点，但是思路是基本上都是这样的，用两个二分查找边界
func searchRange(nums []int, target int) []int {
	if len(nums) == 1 && target == nums[0] {
		return []int{0, 0}
	}
	res := make([]int, 0)
	left1, left2 := 0, 0
	right1, right2 := len(nums)-1, len(nums)-1
	for left1 <= right1 {
		mid := left1 + (right1-left1)/2
		if nums[mid] == target {
			if mid > 0 && nums[mid-1] != target {
				res = append(res, mid)
			} else if mid == 0 {
				res = append(res, mid)
			}
			right1 = mid - 1
		} else if nums[mid] < target {
			left1 = mid + 1
		} else {
			right1 = mid - 1
		}
	}
	for left2 <= right2 {
		mid := left2 + (right2-left2)/2
		if nums[mid] == target {
			if mid < len(nums)-1 && nums[mid+1] != target {
				res = append(res, mid)
			} else if mid == len(nums)-1 {
				res = append(res, mid)
			}
			left2 = mid + 1
		} else if nums[mid] < target {
			left2 = mid + 1
		} else {
			right2 = mid - 1
		}
	}
	if len(res) == 0 {
		return []int{-1, -1}
	}
	return res
}
