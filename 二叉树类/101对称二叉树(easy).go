package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	arr := "[1,2,2,3,4,4,3]"
	root := createTree(arr)
	fmt.Print(isSymmetric_iteration(root))
}

//递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return ismirror(root, root)
}
func ismirror(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return ismirror(a.Left, b.Right) && ismirror(a.Right, b.Left)
	}
	return false
}

//迭代

func isSymmetric_iteration(root *TreeNode) bool {
	if root == nil {
		return true
	}
	L := list.New()
	L.PushBack(root)
	for L.Len() != 0 {
		length := L.Len()
		templist := make([]string, L.Len()) //利用一个临时的string切片记录本层的数据，然后回文串判断
		for i := 0; i < length; i++ {
			temp := L.Remove(L.Front()).(*TreeNode)
			if temp != nil {
				L.PushBack(temp.Left)
				L.PushBack(temp.Right)
				templist[i] = strconv.Itoa(temp.Val)
			} else {
				templist[i] = "null"
			}
		}
		for i := 0; i < length/2; i++ {
			if templist[i] != templist[length-1-i] {
				return false
			}

		}
	}
	return true
}
