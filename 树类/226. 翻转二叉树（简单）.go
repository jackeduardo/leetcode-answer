package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[4,2,7,1,3,6,9]"
	root := Methods.CreateTree(arr)
	Methods.PrintTree(invertTree(root))
}

//非常简单的dfs翻转，只要熟悉递归的套路，做一个左右swap，就可以实现
func invertTree(root *Types.TreeNode) *Types.TreeNode {
	if root == nil {
		return nil
	}
	invertTree_swap(root)
	return root
}

func invertTree_swap(root *Types.TreeNode) {
	if root == nil {
		return
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
