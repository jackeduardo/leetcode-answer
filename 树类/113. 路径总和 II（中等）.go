package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr:="[5,4,8,11,null,13,4,7,2,null,null,5,1]"
	root:=Methods.CreateTree(arr)
	sum:=22
	fmt.Println(pathSum2(root,sum))
}

func pathSum2(root * Types.TreeNode, sum int) [][]int {
	var res [][]int
	var temp []int
	pathSum2_helper(root,sum,&res,&temp)
	return res
}
// 类似回溯，一层函数的功能，再找递归位置
func pathSum2_helper(root * Types.TreeNode, sum int, res*[][]int,temp *[]int){
	if root==nil{
		return
	}
	sum-=root.Val
	*temp=append(*temp,root.Val)
	if root.Left==nil&&root.Right==nil{
		if sum==0{
			out:=make([]int,len(*temp))
			copy(out,*temp)
			*res=append(*res,out)
		}
	}
	pathSum2_helper(root.Left,sum,res,temp)
	pathSum2_helper(root.Right,sum,res,temp)
	*temp=(*temp)[:len(*temp)-1]
}