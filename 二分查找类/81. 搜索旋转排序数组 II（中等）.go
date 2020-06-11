package main

import "fmt"

func main() {
	nums:=[]int{2,5,6,0,0,1,2}
	fmt.Println(search(nums,0))
}

func search81(nums []int, target int) bool {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] { //描述的是target的区间
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] > nums[right]{
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}else {
			//去重，去掉干扰位置,一次次去保证没有遗漏
			if nums[mid]==nums[left]{
				left++
			}
			right--
		}
	}
	return false
}
