package main

import "fmt"

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

//暴力转四角
func rotate(matrix [][]int) {
	n := len(matrix)
	a, b, c, d := []int{0, 0}, []int{0, n - 1}, []int{n - 1, 0}, []int{n - 1, n - 1}
	for a[1] <= b[1] {
		for i := 0; i < b[1]-a[1]; i++ {
			tl := matrix[a[0]][a[1]+i]
			tr := matrix[b[0]+i][b[1]]
			dr := matrix[d[0]][d[1]-i]
			dl := matrix[c[0]-i][c[1]]
			matrix[a[0]][a[1]+i] = dl
			matrix[b[0]+i][b[1]] = tl
			matrix[d[0]][d[1]-i] = tr
			matrix[c[0]-i][c[1]] = dr
		}
		a[0]++
		a[1]++
		b[0]++
		b[1]--
		c[0]--
		c[1]++
		d[0]--
		d[1]--
	}
}

//翻转法，摘抄
func rotate2(matrix [][]int) {
	length := len(matrix)

	// 先左<->右互换
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-1-j] = matrix[i][length-1-j], matrix[i][j]
		}
	}

	// 后关于左下/右上的对角线互换
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			matrix[i][j], matrix[length-j-1][length-i-1] = matrix[length-j-1][length-i-1], matrix[i][j]
		}
	}
}
