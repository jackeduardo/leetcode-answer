package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}

//如果中间的数小于最右边的数，则右半段是有序的，若中间数大于最右边数，则左半段是有序的
func search(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

//递归版
func search2(nums []int, target int) int {
	return search2_helper(nums, 0, len(nums)-1, target)
}
func search2_helper(nums []int, low int, high int, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] < nums[high] {
		if nums[mid] < target && target <= nums[high] {
			return search2_helper(nums, mid+1, high, target)
		} else {
			return search2_helper(nums, low, mid-1, target)
		}
	} else {
		if nums[low] <= target && target < nums[mid] {
			return search2_helper(nums, low, mid-1, target)
		} else {
			return search2_helper(nums, mid+1, high, target)
		}
	}
}
