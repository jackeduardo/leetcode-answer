package main

import "fmt"

func main() {
	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}
	wallsAndGates(rooms)
	fmt.Println(rooms)
}

//从门开始起步的优化bfs
func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}

	DIRECTIONS := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	rLen, cLen := len(rooms), len(rooms[0])
	var queue [][]int
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for len(queue) != 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, direct := range DIRECTIONS {
			r, c := row+direct[0], col+direct[1]
			// 如果撞墙或者访问下一个房间的步数小于等于当前房间步数+1（可能是撞墙了、碰到下一个门、折返了）
			if r < 0 || r >= rLen || c < 0 || c >= cLen || rooms[r][c] <= rooms[row][col]+1 {
				continue
			}

			rooms[r][c] = rooms[row][col] + 1
			queue = append(queue, []int{r, c})
		}
	}

}

//暴力超时bfs
func wallsAndGates2(rooms [][]int) {
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == 2147483647 {
				rooms[i][j] = wallsAndGates_bfs(rooms, i, j)
			}
		}
	}
}

func wallsAndGates_bfs(rooms [][]int, i, j int) int {
	var queue [][]int
	distance := 0
	queue = append(queue, []int{i, j})
	visited := make([][]bool, len(rooms))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(rooms[0]))
	}
	for len(queue) != 0 {
		length := len(queue)
		for p := 0; p < length; p++ {
			pos := queue[0]
			row, col := pos[0], pos[1]
			visited[row][col] = true

			if row >= 1 {
				if rooms[row-1][col] == 0 {
					return distance + 1
				}
				if rooms[row-1][col] != -1 && !visited[row-1][col] {
					queue = append(queue, []int{row - 1, col})
				}
			}
			if row < len(rooms)-1 {
				if rooms[row+1][col] == 0 {
					return distance + 1
				}
				if rooms[row+1][col] != -1 && !visited[row+1][col] {
					queue = append(queue, []int{row + 1, col})
				}
			}
			if col >= 1 {
				if rooms[row][col-1] == 0 {
					return distance + 1
				}
				if rooms[row][col-1] != -1 && !visited[row][col-1] {
					queue = append(queue, []int{row, col - 1})
				}
			}
			if col < len(rooms[0])-1 {
				if rooms[row][col+1] == 0 {
					return distance + 1
				}

				if rooms[row][col+1] != -1 && !visited[row][col+1] {
					queue = append(queue, []int{row, col + 1})
				}
			}
			queue = queue[1:]
		}
		distance++
	}
	return 2147483647
}
