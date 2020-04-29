package main

import (
	"container/list"
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[3,9,20,null,null,15,7],"
	root := Methods.CreateTree(arr)
	fmt.Println(zigzagLevelOrder(root))
}

//层次遍历模板，加上一个h判断方向即可，用curindex作掉头工作
func zigzagLevelOrder(root *Types.TreeNode) [][]int {
	var res [][]int
	var l list.List
	if root == nil {
		return res
	}
	l.PushBack(root)
	h := 1
	for l.Len() > 0 {
		curlen := l.Len()
		curarr := make([]int, curlen)
		for i := 0; i < curlen; i++ {
			node := l.Remove(l.Front()).(*Types.TreeNode)
			curindex := curlen - i - 1
			if h%2 != 0 {
				curindex = i
			}
			curarr[curindex] = node.Val
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		res = append(res, curarr)
		h++
	}
	return res
}
