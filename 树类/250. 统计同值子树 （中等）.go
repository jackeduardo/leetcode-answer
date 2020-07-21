package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[5,1,3,1,1,1,3]"
	root := Methods.CreateTree(arr)
	fmt.Println(countUnivalSubtrees(root))
}

//题目定义没看清楚搞了很久，其实不难，简单dfs
func countUnivalSubtrees(root *Types.TreeNode) int {
	if root == nil {
		return 0
	}
	cur := 0
	if UnivalSubtrees(root) {
		cur++
	}
	return cur + countUnivalSubtrees(root.Left) + countUnivalSubtrees(root.Right)
}

func UnivalSubtrees(root *Types.TreeNode) bool {
	if root == nil {
		return true
	}
	cur := true
	if root.Left != nil && root.Right != nil {
		cur = root.Val == root.Left.Val && root.Val == root.Right.Val
	} else if root.Left != nil {
		cur = root.Val == root.Left.Val
	} else if root.Right != nil {
		cur = root.Val == root.Right.Val
	}
	return cur && UnivalSubtrees(root.Left) && UnivalSubtrees(root.Right)
}
