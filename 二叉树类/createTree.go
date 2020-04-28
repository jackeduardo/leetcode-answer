package main

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(input string )  *TreeNode{
	str:=input[1:len(input)-1]
	strlist:=strings.Split(str,",")
	root:=new(TreeNode)
	res:=new(TreeNode)
	TreeNodeList:=list.New()
	for i := 0; i <len(strlist); i++ {
		if i==0{
			root.Val, _ =strconv.Atoi(strlist[i])
			res=root
			TreeNodeList.PushBack(root)
		}
		if TreeNodeList.Len()!=0{
			root=TreeNodeList.Remove(TreeNodeList.Front()).(*TreeNode)
		}else{
			break
		}
		if i+1<len(strlist)&&strlist[i+1]!="null"{
			value,_:=strconv.Atoi(strlist[i+1])
			root.Left=&TreeNode{Val: value}
			TreeNodeList.PushBack(root.Left)
		}
		if i+2<len(strlist)&&strlist[i+2]!="null"{
			value,_:=strconv.Atoi(strlist[i+2])
			root.Right=&TreeNode{Val: value}
			TreeNodeList.PushBack(root.Right)
		}
		i++
	}
	return res
}