package main

import (
	"fmt"
)

// 解题思路：回溯装填，需要注意的技巧就是，由于题目要求每层都可以反复取，所以递归时传level不要加一，退栈的时候，下个循环会自动加一
func main() {
	candidates := []int{2, 3, 6, 7}
	fmt.Print(combinationSum(candidates, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var temp []int
	if candidates == nil {
		return res
	}
	backtrack_combinationSum(candidates, 0, &res, &temp, target)
	return res
}

func backtrack_combinationSum(candidates []int, level int, res *[][]int, temp *[]int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		out := make([]int, len(*temp))
		copy(out, *temp)
		*res = append(*res, out)
		return
	}
	for ; level < len(candidates); level++ {
		*temp = append(*temp, candidates[level])
		backtrack_combinationSum(candidates, level, res, temp, target-candidates[level])
		*temp = (*temp)[:len(*temp)-1]
	}
}
