package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr:="[1,2,3,4,5,6]"
	root:=Methods.CreateTree(arr)
	fmt.Println(countNodes(root))
}
//通俗暴力递归解，遍历一遍即可
func countNodes(root *Types.TreeNode) int {
	res:=0
	countNodes_helper(root,&res)
	return res
}
func countNodes_helper(root *Types.TreeNode, res *int)  {
	if root==nil{
		return
	}
	*res+=1
	countNodes_helper(root.Left,res)
	countNodes_helper(root.Right,res)
}