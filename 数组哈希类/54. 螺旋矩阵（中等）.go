package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3,4},
		{5, 6, 7,8},
		{9, 10, 11,12},
	}
	fmt.Println(spiralOrder(matrix))
}
//螺旋搜索，由于是矩阵而不一定是正方形，所以在循环内部添加边界判断，不符合就break，否则如果是长方形，那么会多出来不必要的数据
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return nil
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res

}
