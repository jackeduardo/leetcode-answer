package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr:="[1,2,3,null,5,null,4]"
	root:=Methods.CreateTree(arr)
	fmt.Println(rightSideView(root))
}
//使用bfs层序遍历，并只保留每层最后一个节点的值
func rightSideView(root *Types.TreeNode) []int {
	if root==nil{
		return nil
	}
	var res []int
	var queue []*Types.TreeNode
	queue=append(queue,root)
	for len(queue)!=0{
		length:=len(queue)
		for i := 0; i <length; i++ {
			node:=queue[0]
			if i==length-1{
				res = append(res, node.Val)
			}
			if node.Left!=nil{
				queue = append(queue, node.Left)
			}
			if node.Right!=nil{
				queue = append(queue, node.Right)
			}
			queue=queue[1:]
		}
	}
	return res
}