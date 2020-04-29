package Methods

import (
	"container/list"
	"leetcode-answer/Types"
	"strconv"
	"strings"
)

func CreateTree(input string) *Types.TreeNode {
	str := input[1 : len(input)-1]
	strlist := strings.Split(str, ",")
	root := new(Types.TreeNode)
	res := new(Types.TreeNode)
	TreeNodeList := list.New()
	for i := 0; i < len(strlist); i++ {
		if i == 0 {
			root.Val, _ = strconv.Atoi(strlist[i])
			res = root
			TreeNodeList.PushBack(root)
		}
		if TreeNodeList.Len() != 0 {
			root = TreeNodeList.Remove(TreeNodeList.Front()).(*Types.TreeNode)
		} else {
			break
		}
		if i+1 < len(strlist) && strlist[i+1] != "null" {
			value, _ := strconv.Atoi(strlist[i+1])
			root.Left = &Types.TreeNode{Val: value}
			TreeNodeList.PushBack(root.Left)
		}
		if i+2 < len(strlist) && strlist[i+2] != "null" {
			value, _ := strconv.Atoi(strlist[i+2])
			root.Right = &Types.TreeNode{Val: value}
			TreeNodeList.PushBack(root.Right)
		}
		i++
	}
	return res
}
