package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCB"))
}
//依然是回溯算法的应用
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if exist_dfs(board, word, i, j, visited, 0) {
				return true
			}
		}
	}
	return false
}

func exist_dfs(board [][]byte, word string, i int, j int, visited [][]bool, index int) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 {
		return false
	}
	if word[index] != board[i][j] {
		return false
	}
	if !visited[i][j] {
		visited[i][j] = true
		if exist_dfs(board, word, i-1, j, visited, index+1) {
			return true
		}
		if exist_dfs(board, word, i+1, j, visited, index+1) {
			return true
		}
		if exist_dfs(board, word, i, j-1, visited, index+1) {
			return true
		}
		if exist_dfs(board, word, i, j+1, visited, index+1) {
			return true
		}
		visited[i][j] = false
	}
	return false
}
