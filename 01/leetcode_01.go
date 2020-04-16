package main

import (
"container/list"
"fmt"

)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val       int
	Neighbors []*Node
}
type Pos struct {
	i int
	j int
}
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
var pc [256]byte

func main() {
	name:=make([]string,0,2)

	fmt.Println(name)
}
func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]

}
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func gcd(x,y int) int{
	for y!=0{
		x,y=y,x%y
	}
	return x
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || i == m-1 || j == n-1
			if isEdge && board[i][j] == 'O' {
				bfs(board, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

}

func bfs(board [][]byte, i int, j int) {
	var queue list.List
	queue.PushBack(Pos{i, j})
	board[i][j] = '#'
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*Pos)
		if cur.i-1 >= 0 && board[cur.i-1][cur.j] == 'O' {
			queue.PushBack(Pos{cur.i - 1, cur.j})
			board[cur.i-1][cur.j] = '#'
		}
		if cur.i+1 <= len(board)-1 && board[cur.i+1][cur.j] == 'O' {
			queue.PushBack(Pos{cur.i + 1, cur.j})
			board[cur.i+1][cur.j] = '#'
		}
		if cur.j-1 >= 0 && board[cur.i][cur.j-1] == 'O' {
			queue.PushBack(Pos{cur.i, cur.j - 1})
			board[cur.i][cur.j-1] = '#'
		}
		if cur.j+1 <= len(board[0])-1 && board[cur.i][cur.j+1] == 'O' {
			queue.PushBack(Pos{cur.i, cur.j + 1})
			board[cur.i][cur.j+1] = '#'
		}
	}

}

func cloneGraph(node *Node) *Node {
	if node==nil{
		return nil
	}
	queue:=list.New()
	queue.PushBack(node)
	visited:=map[*Node]*Node{}
	visited[node]=&Node{Val:node.Val,Neighbors:make([]*Node,len(node.Neighbors))}
	for queue.Len()>0{
		node:=queue.Remove(queue.Front()).(*Node)
		for i,v:=range node.Neighbors{
			if _,ok:=visited[v];!ok{
				queue.PushBack(v)
				visited[v]=&Node{v.Val,make([]*Node,len(v.Neighbors))}

			}
			visited[node].Neighbors[i]=visited[v]
		}
	}
	return visited[node]

}
