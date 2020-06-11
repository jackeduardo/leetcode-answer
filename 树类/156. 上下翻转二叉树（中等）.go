package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	root:=Methods.CreateTree("[1,2]")
	Methods.PrintTree(upsideDownBinaryTree(root))
}

//自己实现的方案，用了一个变量来记录新的root
func upsideDownBinaryTree(root *Types.TreeNode) *Types.TreeNode {
	newroot:=root
	upsideDownBinaryTree_helper(&newroot,root)
	return newroot
}

func upsideDownBinaryTree_helper(newroot **Types.TreeNode,root *Types.TreeNode) *Types.TreeNode {
	if root==nil{
		return root
	}
	if root.Left!=nil&& root.Left.Left==nil{
		*newroot=root.Left
	}
	left:=upsideDownBinaryTree_helper(newroot,root.Left)
	right:=upsideDownBinaryTree_helper(newroot,root.Right)
	if left!=nil{
		left.Left=right
		left.Right=root
		root.Left=nil
		root.Right=nil
	}
	return root
}
//摘抄的递归方案

func upsideDownBinaryTree2(root *Types.TreeNode) *Types.TreeNode {
	if root==nil||root.Left==nil{
		return root
	}
	left:=root.Left
	right:=root.Right
	node:=upsideDownBinaryTree2(root.Left)//因为本题中所有的右节点要么具有兄弟节点，要么为空，所以新的root必然只在左枝上
	root.Left=nil
	root.Right=nil
	left.Left=right
	left.Right=root
	return node
}