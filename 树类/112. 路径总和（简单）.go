package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[5,4,8,11,null,13,4,7,2,null,null,null,1]"
	root := Methods.CreateTree(arr)
	sum := 22
	fmt.Println(hasPathSum(root, sum))
}

func hasPathSum(root *Types.TreeNode, sum int) bool {
	if root == nil {
		return false // 这个判断语句的实际抽象逻辑是，到空节点的路径不符合题意，可以理解为（不存在到达空节点的路径->没有路径->路径合为空）所以返回false，由于函数的主要判断是||，这里不会影响结果
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)

}
