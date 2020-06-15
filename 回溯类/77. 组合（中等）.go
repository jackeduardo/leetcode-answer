package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

//标准回溯算法组合模板
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	combine_backtrack(&res, k, arr, 0, make([]int, 0))
	return res

}

func combine_backtrack(res *[][]int, k int, arr []int, index int, cur []int) {
	if k == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
	}
	for i := index; i < len(arr); i++ {
		cur = append(cur, arr[i])
		combine_backtrack(res, k-1, arr, i+1, cur)
		cur = cur[:len(cur)-1]
	}

}
