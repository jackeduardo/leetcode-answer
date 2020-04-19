package main

import "fmt"
//思路：关键是利用回溯算法的特性：能够按照条件来遍历所有解，不符合就退，符合就前进，本质是dfs的搜索
func main() {
	board:=[][]byte{
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
	fmt.Println("Previous:")
	PrintBoard(board)
	solveSudoku(board)
	fmt.Println("After:")
	PrintBoard(board)
}

func PrintBoard(board [][]byte)  {
	for i := 0; i <len(board); i++ {
		for j := 0; j <len(board[0]); j++ {
			fmt.Print(string(board[i][j])," ")
		}
		fmt.Println()
	}
}

func solveSudoku(board [][]byte)  {
	row:=[9][9]bool{}
	col:=[9][9]bool{}
	block:=[9][9]bool{}
	for i := 0; i <9; i++ {
		for j := 0; j <9; j++ {
			if board[i][j]!='.'{
				num:=board[i][j]-'1'
				row[i][num]=true
				col[j][num]=true
				block[i/3*3+j/3][num]=true// 这里是取格子的索引，务必记一下，自己推还是有点麻烦，或许有相关规律可一探索，以后再说吧
			}
		}
	}
	backtrackSoduku(board,row,col,block,0,0)
}
//由于每一步（填充）的前进都要判断是否符合条件，并且影响着前一步是否该放弃，所以此函数返回值应该是bool，void的方法我还没有想出来
func backtrackSoduku(board [][]byte, row, col, block [9][9]bool, i, j int) bool {
	//边界校验，为了换行需求，之后如果填充完成，return true
	if j==len(board[0]){
		j=0
		i++
		if i==len(board) {
			return true
		}
	}
	if board[i][j]=='.' {
		for n := 0; n<9; n++ {
			blockindex:=i/3*3+j/3
			if !row[i][n]&&!col[j][n]&&!block[blockindex][n] {
				board[i][j]=byte('1'+n)
				row[i][n]=true
				col[j][n]=true
				block[blockindex][n]=true
				if backtrackSoduku(board,row,col,block,i,j+1){
					return true
				}
				board[i][j]='.'
				row[i][n]=false
				col[j][n]=false
				block[blockindex][n]=false
			}
		}
	}else{
		return backtrackSoduku(board,row,col,block,i,j+1)
	}
	return false
}