package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[3,9,20,null,null,15,7]"
	root := Methods.CreateTree(arr)
	fmt.Println(maxDepth2(root))
}

// dfs分治到最深处，回退一层加一，通过max函数维护最深的数值，速度较慢
func maxDepth(root *Types.TreeNode) int {
	if root == nil {
		return 0
	}
	return Methods.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//dfs，打标记的常规办法
func maxDepth2(root *Types.TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	dfs(root, 0, &res)
	return res
}

func dfs(root *Types.TreeNode, level int, res *int) {
	if root == nil {
		return
	}
	if *res < level+1 {
		*res = level + 1
	}
	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)
}
