package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

//无返回值dfs搜索，每次搜索到0就返回，那么一共调用的搜索次数就是岛屿的数量
func numIslands(grid [][]byte) int {
	count := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				numIslands_dfs(grid, i, j, visited)
				count++
			}
		}
	}
	return count
}

func numIslands_dfs(grid [][]byte, i, j int, visited [][]bool) {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == '0' {
		return
	}
	if !visited[i][j] {
		visited[i][j] = true
		numIslands_dfs(grid, i-1, j, visited)
		numIslands_dfs(grid, i, j+1, visited)
		numIslands_dfs(grid, i+1, j, visited)
		numIslands_dfs(grid, i, j-1, visited)
	}
}
