package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {

	arr:="[-10,9,20,null,null,15,7]"
	root:=Methods.CreateTree(arr)
	fmt.Println(maxPathSum(root))
}
func maxPathSum(root *Types.TreeNode) int {
	var res=-int(^uint(0) >> 1)
	maxPathSum_getmax(root,&res)
	return res
}
//难点在于递归中左右子树只能取一边较大的，因为取两边就不符合路径要求，但是题目要求的最大路径合是有可能左子树加上右子树和节点的，所以最大值res里面是需要更新left+right+root.val
func maxPathSum_getmax(root *Types.TreeNode,res *int)int{
	if root==nil{
		return 0
	}
	left:=Methods.Max(0,maxPathSum_getmax(root.Left,res))//左子树最大路径合
	right:=Methods.Max(0,maxPathSum_getmax(root.Right,res))//右子树最大路径合
	*res=Methods.Max(left+right+root.Val,*res)
	return Methods.Max(left,right)+root.Val
}
