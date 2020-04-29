package main

import (
	"container/list"
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

//bfs层序遍历 模板题
func main() {
	arr := "[3,9,20,null,null,15,7],"
	root := Methods.CreateTree(arr)
	fmt.Println(levelOrder(root))
}

func levelOrder(root *Types.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var l list.List
	l.PushBack(root)
	for l.Len() > 0 {
		curlen := l.Len()
		curarr := make([]int, curlen)
		for i := 0; i < curlen; i++ {
			node := l.Remove(l.Front()).(*Types.TreeNode)
			curarr[i] = node.Val
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		res = append(res, curarr)
	}
	return res
}
