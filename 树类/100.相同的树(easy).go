package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr1 := "[1,2,3]"
	arr2 := "[1,2,3]"
	root1 := Methods.CreateTree(arr1)
	root2 := Methods.CreateTree(arr2)
	fmt.Print(isSameTree(root1, root2))
}

func isSameTree(p *Types.TreeNode, q *Types.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
