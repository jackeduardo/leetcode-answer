package main

import (
	"container/list"
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := "[3,9,20,null,null,15,7]"
	root := Methods.CreateTree(arr)
	fmt.Println(levelOrderBottom(root))
}
func levelOrderBottom(root *Types.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	List := list.New()
	List.PushBack(root)
	for List.Len() != 0 {
		length := List.Len()
		templist := make([]int, length)
		for i := 0; i < length; i++ {
			temproot := List.Remove(List.Front()).(*Types.TreeNode)
			if temproot.Left != nil {
				List.PushBack(temproot.Left)
			}
			if temproot.Right != nil {
				List.PushBack(temproot.Right)
			}
			templist[i] = temproot.Val
		}
		res = append([][]int{templist}, res...)
	}
	return res
}