package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[2,1,3]"
	root := Methods.CreateTree(arr)
	fmt.Println(isValidBST2(root))
}

//设置pre的中序遍历的普通方法
func isValidBST(root *Types.TreeNode) bool {
	var pre_isValidBST = &Types.TreeNode{Val: -int(^uint(0) >> 1)}
	var res = true
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	isValidBST_helper(root, &res, &pre_isValidBST)
	return res

}

func isValidBST_helper(root *Types.TreeNode, res *bool, pre **Types.TreeNode) {
	if root == nil {
		return
	}
	isValidBST_helper(root.Left, res, pre)
	(*pre).Right = root
	if (*pre).Val < root.Val {
		*res = *res && true
	} else {
		*res = *res && false
	}
	*pre = (*pre).Right
	isValidBST_helper(root.Right, res, pre)
}

//另一种写法
var last *Types.TreeNode

func isValidBST2(root *Types.TreeNode) bool {
	last = &Types.TreeNode{Val: -1 << 63}
	return isValidBST2_helper(root)
}

func isValidBST2_helper(root *Types.TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST2_helper(root.Left) || root.Val <= last.Val {
		return false
	}
	last = root
	return isValidBST2_helper(root.Right)
}
