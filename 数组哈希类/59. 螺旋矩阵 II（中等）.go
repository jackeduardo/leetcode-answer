package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

//由于是正方形，不必担心会多出来数据，因为正方形中一个边界条件不符合，其他边界条件肯定都不符合，不会进入不必要的循环
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	val := 1
	top, down, left, right := 0, n-1, 0, n-1
	for val <= n*n {
		for i := left; i <= right; i++ {
			res[top][i] = val
			val++
		}
		top++

		for i := top; i <= down; i++ {
			res[i][right] = val
			val++
		}
		right--

		for i := right; i >= left; i-- {
			res[down][i] = val
			val++
		}
		down--

		for i := down; i >= top; i-- {
			res[i][left] = val
			val++
		}
		left++

	}
	return res
}
