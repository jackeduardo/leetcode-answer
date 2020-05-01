package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[3,9,20,null,null,15,7]"
	root := Methods.CreateTree(arr)
	fmt.Println(isBalanced(root))
}

func isBalanced(root *Types.TreeNode) bool {
	return isBalanced_height(root)!=-1
}
//从底至顶，不太容易想，感觉理解就行
func isBalanced_height(root *Types.TreeNode) int{
	if root==nil{
		return 0
	}
	left:=isBalanced_height(root.Left)
	if left==-1{
		return -1
	}
	right:=isBalanced_height(root.Right)
	if right==-1{
		return -1
	}
	if Methods.Abs(left-right)<2{//高度差判断，返回最大高度
		return Methods.Max(left,right)+1
	}
	return -1
}

//从顶至底（暴力法）
func isBalanced2(root *Types.TreeNode) bool {
	if root==nil{
		return true
	}
	return Methods.Abs(isBalanced_depth(root.Left)-isBalanced_depth(root.Right))<=1&&isBalanced2(root.Left)&&isBalanced2(root.Right)//对于当前的节点的判断和它左右节点的递归
}
//获得最大深度
func isBalanced_depth(root *Types.TreeNode) int {
	if root==nil{
		return 0
	}
	return Methods.Max(isBalanced_depth(root.Left),isBalanced_depth(root.Right))+1
}