package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	root:=buildTree_106(inorder,postorder)
	Methods.PrintTree(root)
}
//分治递归，按照中序和后序的关系，推断出每次截取的范围来确定左右子树,递归到最深处返回，每上一层和当前层串联，是一个回溯的过程，所以分治就是找到最（深）小单位，然后回溯得到结果
//优化了一下，减少复杂度
func buildTree_106(inorder []int, postorder []int) *Types.TreeNode {
	if len(inorder)==0||len(postorder)==0{
		return nil
	}
	for i, v := range inorder {
		if v==postorder[len(postorder)-1]{
			return &Types.TreeNode{
				Val:   inorder[i],
				Left:  buildTree_106(inorder[0:i],postorder[0:i]),
				Right: buildTree_106(inorder[i+1:],postorder[i:len(postorder)-1]),
			}
		}
	}
	return nil
}
