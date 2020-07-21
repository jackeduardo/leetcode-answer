package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix2_240(matrix, 5))
}

//对角线法
func searchMatrix_240(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

//逐行二分，利用矩阵特性的优化二分
func searchMatrix2_240(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	l, r, mid := 0, len(matrix[0])-1, 0
	for i := 0; i < len(matrix); i++ {
		for l <= r {
			mid = l + (r-l)/2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		l = 0
	}
	return false
}
