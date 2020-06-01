package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

//利用二维数组来查重
func isValidSudoku(board [][]byte) bool {
	row := [9][9]bool{}
	col := [9][9]bool{}
	block := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				cur := int(board[i][j]-'0') - 1
				if !row[i][cur] && !col[j][cur] && !block[i/3*3+j/3][cur] {
					row[i][cur] = true
					col[j][cur] = true
					block[i/3*3+j/3][cur] = true
				} else {
					return false
				}
			}
		}
	}
	return true
}
