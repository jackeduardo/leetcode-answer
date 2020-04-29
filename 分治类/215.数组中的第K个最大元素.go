package main

import (
	"fmt"
	"math/rand"
	"time"
)

//最简单的方法就是调用sort库函数，但是这里是训练使用分治法
//分治的核心就是找到最小单位进行按需计算
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Print(findKthLargest(nums, 2))

}

//核心思路就是给随机的pivot排到它应该所处的序列位置，例如pivot如果第三个最大，那它就在倒数第三个位置，完成这个操作就是通过partition
func findKthLargest(nums []int, k int) int {
	len := len(nums)
	left := 0
	right := len - 1
	target := len - k
	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func partition(nums []int, left int, right int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(right+1-left) + left
	nums[left], nums[randomIndex] = nums[randomIndex], nums[left]
	pivot := nums[left]
	lt := left + 1
	rt := right
	for lt <= rt {
		for lt <= rt && nums[lt] < pivot {
			lt++
		}
		for lt <= rt && nums[rt] > pivot {
			rt--
		}

		if lt >= rt {
			break
		}
		nums[lt], nums[rt] = nums[rt], nums[lt]
		lt++
		rt--
	}
	nums[left], nums[rt] = nums[rt], nums[left]
	return rt
}
