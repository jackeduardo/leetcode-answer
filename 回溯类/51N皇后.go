package main

import "fmt"

//var res [][]string

func main() {
	res:=solveNQueens(4)
	printQueens(res)
}
func printQueens(string[][]string){
	for _, r := range string {
		for _, c := range r {
			fmt.Print(c," ")
		}
		fmt.Println()
	}
}

func solveNQueens(n int) [][]string {
	row:=make([]byte,n)
	board:=make([]string,n)
	var res [][]string
	for  i := range row {
		row[i]='.'
	}
	for i:=range board{
		board[i]=string(row)
	}
	backtrack_solveNQueens(board,0,&res)
	return res
}

func backtrack_solveNQueens(board []string,row int,res*[][]string){
	if row==len(board){
		out:=make([]string,len(board))
		copy(out,board)
		*res=append(*res,out)
		return
	}
	n:=len(board[row])
	for col := 0; col<n; col++ {
		if !isvaild(board,row,col){
			continue
		}
		temp:=[]byte(board[row])
		temp[col]='Q'
		board[row]=string(temp)
		backtrack_solveNQueens(board,row+1,res)
		temp[col]='.'
		board[row]=string(temp)
	}
}
//判断是否可以添加Q，其实可以写8个方向，但是由于一行一行搜索判断，所以左下右下本行的方向都不用判断
func isvaild(board []string,row int,col int) bool{
	n:=len(board)
	for i := 0; i <n; i++ {
		if board[i][col]=='Q'{
			return false
		}
	}
	for i,j:= row-1,col+1; i>=0&&j<n; i,j=i-1,j+1 {
		if board[i][j]=='Q'{
			return false
		}
	}
	for i,j:= row-1,col-1; i>=0&&j>=0; i,j=i-1,j-1 {
		if board[i][j]=='Q'{
			return false
		}
	}
	return true
}