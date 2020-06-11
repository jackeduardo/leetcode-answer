package main

import "fmt"

func main() {
	board:=[][]byte{
		{'X','X','X','X'},
		{'X','O','O','X'},
		{'X','X','O','X'},
		{'X','O','X','X'},
	}
	solve(board)
	fmt.Println(board)
}
func solve(board [][]byte)  {
	if board==nil||len(board)==0{
		return
	}
	m:=len(board)
	n:=len(board[0])
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			isEdge:=i==0||i==m-1||j==0||j==n-1
			if isEdge&&board[i][j]=='O'{
				dfs(board,i,j)
			}
		}
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if board[i][j]=='O'{
				board[i][j]='X'
			}
			if board[i][j]=='#'{
				board[i][j]='O'
			}

		}
	}
}

func dfs(board[][] byte, i int ,j int){
	if i<0||i>=len(board)||j<0||j>=len(board[0])||board[i][j]=='X'||board[i][j]=='#'{
		return
	}
	board[i][j]='#'
	dfs(board,i-1,j)
	dfs(board,i+1,j)
	dfs(board,i,j-1)
	dfs(board,i,j+1)
}