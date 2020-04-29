package main

import "fmt"

// 解题思路，回溯减枝
func main() {
	nums := []int{1, 2, 3}
	fmt.Print(subsets(nums))
}

func subsets(nums []int) [][]int {
	var res [][]int
	var temp []int
	visited := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}
	backtrack_subsets(nums, 0, &temp, len(nums), &res, visited)
	res = append(res, make([]int, 0))
	return res
}

func backtrack_subsets(nums []int, level int, temp *[]int, length int, res *[][]int, visited []int) {
	if len(*temp) != 0 {
		out := make([]int, len(*temp))
		copy(out, *temp)
		*res = append(*res, out)
	}
	for i := level; i < length; i++ {
		if visited[i] == 1 {
			continue
		}
		visited[i] = 1
		*temp = append(*temp, nums[i])
		backtrack_subsets(nums, i, temp, length, res, visited) //注意这里i的位置不能是level+1，否则的话回溯到第一层的时候又会多递归两次
		*temp = (*temp)[:len(*temp)-1]
		visited[i] = 0
	}

}
