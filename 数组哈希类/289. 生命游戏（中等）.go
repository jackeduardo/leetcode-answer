package main

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(board)
}

//辅助数组的通用方法,可以通过bit操作来优化掉辅助数组（见leetcode的题解）
func gameOfLife(board [][]int) {
	if board == nil || board[0] == nil {
		return
	}
	temp := make([][]int, len(board))
	for i := 0; i < len(temp); i++ {
		temp[i] = make([]int, len(board[0]))
		copy(temp[i], board[i])
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			changestate(temp, board, i, j)
		}
	}

}

func changestate(temp [][]int, board [][]int, i, j int) {
	directions := [][]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 0},
		{0, -1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
	count := 0
	for _, v := range directions {
		row, col := i+v[0], j+v[1]
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0])-1 {
			continue
		}
		if temp[row][col] == 1 {
			count++
		}
	}
	if temp[i][j] == 1 {
		if count < 2 {
			board[i][j] = 0
		}
		if count == 2 || count == 3 {
			board[i][j] = 1
		}
		if count > 3 {
			board[i][j] = 0
		}
	}
	if temp[i][j] == 0 {
		if count == 3 {
			board[i][j] = 1
		}
	}
}
