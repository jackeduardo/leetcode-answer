package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	Methods.PrintTree(root)
}

//分治递归，按照中序和前序的关系，推断出每次截取的范围来确定左右子树,递归到最深处返回，每上一层和当前层串联，是一个回溯的过程，所以分治就是找到最（深）小单位,回溯得到结果
func buildTree(preorder []int, inorder []int) *Types.TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}
	for i, v := range inorder {
		if v == preorder[0] {
			return &Types.TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:1+i], inorder[0:i]),
				Right: buildTree(preorder[1+i:], inorder[i+1:]),
			}
		}
	}
	return nil

}
