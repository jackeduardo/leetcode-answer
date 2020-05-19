package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2, 2, 2}
	fmt.Println(makesquare(nums))
}

func makesquare(nums []int) bool {
	if nums == nil {
		return false
	}
	if len(nums) < 4 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%4 != 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if sum/4 < nums[i] {
			return false
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) //排序降复杂度
	//return makesquare_dfs(nums, 0, len(nums), 0, 0, 0, 0, sum/4)
	return makesquare_backtrack(nums, 0, []int{0, 0, 0, 0}, sum/4)

}

//dfs，每一层就是每个index，每一层4条边都去尝试构造边长加上num[index]，超长了就返回false，最后四条边里都是保存着最后的长度，如果长度都是side，那么返回true
func makesquare_dfs(nums []int, i int, length int, i1 int, i2 int, i3 int, i4 int, side int) bool {
	if i == length {
		if i1 == side && i2 == side && i3 == side && i4 == side {
			return true
		} else {
			return false
		}
	}
	if i1 > side || i2 > side || i3 > side || i4 > side {
		return false
	}
	side1 := makesquare_dfs(nums, i+1, length, i1+nums[i], i2, i3, i4, side)
	side2 := makesquare_dfs(nums, i+1, length, i1, i2+nums[i], i3, i4, side)
	side3 := makesquare_dfs(nums, i+1, length, i1, i2, i3+nums[i], i4, side)
	side4 := makesquare_dfs(nums, i+1, length, i1, i2, i3, i4+nums[i], side)
	return side1 || side2 || side3 || side4
}

//回溯剪枝，分别对每一条边构造

func makesquare_backtrack(nums []int, pos int, sides []int, side int) bool {
	if pos == len(nums) {
		if sides[0] == side && sides[1] == side && sides[2] == side && sides[3] == side {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < 4; i++ {
		if sides[i]+nums[pos] > side {
			continue
		}
		sides[i] += nums[pos]
		if makesquare_backtrack(nums, pos+1, sides, side) {
			return true
		}
		sides[i] -= nums[pos]
	}
	return false
}
