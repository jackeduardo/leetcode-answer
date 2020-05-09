package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {

	arr := "[4,2,5,1,3]"
	root := Methods.CreateTree(arr)
	res := treeToDoublyList(root)
	fmt.Println(res.Val)
	//Methods.PrintTree(res)
}

var prenode = &Types.TreeNode{Val: 0}
var tail = &Types.TreeNode{Val: 0}

func treeToDoublyList(root *Types.TreeNode) *Types.TreeNode {
	if root == nil {
		return nil
	}
	head := prenode
	//tail:=&Types.TreeNode{Val:0}
	treeToDoublyList_helper(root)
	head = head.Right
	tail.Right = head
	head.Left = tail
	return head
}

//
func treeToDoublyList_helper(root *Types.TreeNode) {
	if root == nil {
		return
	}
	treeToDoublyList_helper(root.Left)
	prenode.Right = root
	root.Left = prenode
	prenode = prenode.Right
	tail = prenode
	treeToDoublyList_helper(root.Right)

}
