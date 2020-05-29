package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

//双指针，先排序，迭代的时候跳过重复值即可
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
	return res

}

//超时dfs，利用回溯特性再剪枝重复
func threeSum2(nums []int) [][]int {
	var res [][]int
	var cur []int
	sort.Ints(nums)
	threeSum_dfs(nums, &res, cur, 0)
	return res
}

func threeSum_dfs(nums []int, res *[][]int, cur []int, index int) {
	if len(cur) == 3 {
		if cur[0]+cur[1]+cur[2] == 0 {
			copy_slice := make([]int, len(cur))
			copy(copy_slice, cur)
			*res = append(*res, copy_slice)
		}
		return
	}
	temp := math.MaxInt64
	for i := index; i < len(nums); i++ {
		if nums[i] != temp {
			temp = nums[i]
			cur = append(cur, nums[i])
			threeSum_dfs(nums, res, cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
}
