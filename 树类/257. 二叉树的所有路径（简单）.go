package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
	"strconv"
)

func main() {
	arr:="[1,2,3,null,5]"
	root:=Methods.CreateTree(arr)
	fmt.Println(binaryTreePaths(root))
}
//常规dfs
func binaryTreePaths(root *Types.TreeNode) []string {
	res:=make([]string,0)
	binaryTreePaths_helper(root,&res,make([]int,0))
	return res
}

func binaryTreePaths_helper(root *Types.TreeNode,res *[]string,temp []int){
	if root==nil{
		return
	}
	temp=append(temp,root.Val)
	if root.Left==nil&&root.Right==nil{
		*res=append(*res,formString(temp))
	}
	binaryTreePaths_helper(root.Left,res,temp)
	binaryTreePaths_helper(root.Right,res,temp)
}

func formString(temp []int) string {
	s:=""
	arrow:="->"
	for i := 0; i <len(temp); i++ {
		value:=strconv.Itoa(temp[i])
		s+=value+arrow
	}
	return s[:len(s)-2]
}