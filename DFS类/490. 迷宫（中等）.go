package main

import "fmt"

func main() {
	maze := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	start := []int{0, 0}
	destination := []int{1, 2}
	fmt.Println(hasPath(maze, start, destination))
}

//首先这题题意非常模糊，球停下来的条件是：直线行进过程中撞墙，撞墙点去和终点来对比才能判断是否成功
//这种搜索类的题目，递归的要点就是一个递归表示一个搜索，进入下一个搜索的节点，每个函数就是当前节点的功能和操作
//所以这里不符合要求的就不进入递归
//bfs版本之后有空再补上吧
var res bool

func hasPath(maze [][]int, start []int, destination []int) bool {
	res = false
	hasPath_dfs(maze, start[0], start[1], destination)
	return res

}

func hasPath_dfs(maze [][]int, i int, j int, destination []int) {
	if maze[i][j] == 1 || maze[i][j] == -1 {
		return
	}

	maze[i][j] = -1
	top, right, down, left := i-1, j+1, i+1, j-1
	for top >= 0 && maze[top][j] == 0 {
		top--
	}
	if top < 0 || maze[top][j] == 1 {
		res = res || top+1 == destination[0] && j == destination[1]
		hasPath_dfs(maze, top+1, j, destination)
	}
	for right <= len(maze[0])-1 && maze[i][right] == 0 {
		right++
	}
	if right > len(maze[0])-1 || maze[i][right] == 1 {
		res = res || i == destination[0] && right-1 == destination[1]
		hasPath_dfs(maze, i, right-1, destination)
	}
	for down <= len(maze)-1 && maze[down][j] == 0 {
		down++
	}
	if down > len(maze)-1 || maze[down][j] == 1 {
		res = res || down-1 == destination[0] && j == destination[1]
		hasPath_dfs(maze, down-1, j, destination)
	}
	for left >= 0 && maze[i][left] == 0 {
		left--
	}
	if left < 0 || maze[i][left] == 1 {
		res = res || i == destination[0] && left+1 == destination[1]
		hasPath_dfs(maze, i, left+1, destination)
	}

}

//var res bool
//
//func hasPath(maze [][]int, start []int, destination []int) bool {
//	visited := make([][]bool, len(maze))
//	for i, _ := range visited {
//		visited[i] = make([]bool, len(maze[0]))
//	}
//	return hasPath_dfs(maze, start, destination, visited)
//
//}
//
//func hasPath_dfs(maze [][]int, start []int, destination []int, visited [][]bool) bool {
//	if visited[start[0]][start[1]] {
//		return false
//	}
//	if start[0] == destination[0] && start[1] == destination[1] {
//		return true
//	}
//	visited[start[0]][start[1]] = true
//	right, left, up, down := start[1]+1, start[1]-1, start[0]-1, start[0]+1
//	for right < len(maze[0]) && maze[start[0]][right] == 0 {
//		right++
//	}
//	if hasPath_dfs(maze, []int{start[0], right - 1}, destination, visited) {
//		return true
//	}
//	for left >= 0 && maze[start[0]][left] == 0 {
//		left--
//	}
//	if hasPath_dfs(maze, []int{start[0], left + 1}, destination, visited) {
//		return true
//	}
//	for up >= 0 && maze[up][start[1]] == 0 {
//		up--
//	}
//	if hasPath_dfs(maze, []int{up + 1, start[1]}, destination, visited) {
//		return true
//	}
//
//	for down < len(maze) && maze[down][start[1]] == 0 {
//		down++
//	}
//	if hasPath_dfs(maze, []int{down - 1, start[1]}, destination, visited) {
//		return true
//	}
//	return false
//
//}
