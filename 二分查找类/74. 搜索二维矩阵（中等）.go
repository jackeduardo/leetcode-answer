package main

import "fmt"

func main() {
	matrix:=[][]int{
		{1,   3,  5,  7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix,3))
}


func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0{
		return false
	}
	m:=len(matrix)
	n:=len(matrix[0])
	for i := 0; i <m; i++ {
		for j := 0; j <n; j++ {
			if matrix[i][0]<=target&&target<=matrix[i][n-1]{
				return searchMatrix_bs(matrix[i],target)
			}
		}
	}
	return false
}

func searchMatrix_bs(cur []int,target int) bool{
	left:=0
	right:=len(cur)-1
	for left<=right{
		mid:=left+(right-left)/2
		if cur[mid]==target{
			return true
		}else if cur[mid]<target{
			left=mid+1
		}else {
			right=mid-1
		}
	}
	return false
}
//利用升序矩阵的特性，用右上角（判断当前行的最后一个数值）开始作条件判断和遍历
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix)==0{
		return false
	}
	m:=len(matrix)
	n:=len(matrix[0])
	row,col:=0,n-1
	for row<m&&col>=0{
		if matrix[row][col]<target{
			row++
		}else if matrix[row][col]>target{
			col--
		}else {
			return true
		}
	}
	return false
}