package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}

//正解,符合要求的update就行
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

//丑陋的初始代码
func removeElement2(nums []int, val int) int {
	count := 0
	left := 0
	right := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			left = i
			break
		}
		if i == len(nums)-1 {
			return len(nums)
		}
	}
	for i := left; i < len(nums); i++ {
		if nums[i] == val {
			count++
		}
		if i < len(nums)-1 && nums[i+1] != val {
			right = i + 1
			break
		}
		if i == len(nums)-1 && nums[i] == val {
			return len(nums[:len(nums)-count])
		}
	}
	temp := right
	for i := 0; i < len(nums)-temp; i++ {
		nums[left] = nums[right]
		left++
		right++
	}
	return len(nums[:len(nums)-count])
}
