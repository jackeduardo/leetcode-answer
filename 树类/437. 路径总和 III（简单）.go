package main

import (
	"fmt"
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr:="[10,5,-3,3,2,null,11,3,-2,null,1]"
	root:=Methods.CreateTree(arr)
	sum:=8
	fmt.Println(pathSum3(root,sum))
}
//双重递归，没什么难的，就是golang在leetcode上面有坑点全局变量不能随便用，因为golang的全局变量包下共享，所以所有的testcase也会共享，导致提交会通不过，但是单独的测试用例能pass
//所以每次的测试用例需要重置一遍，感觉有点蠢
var res int

func pathSum3(root *Types.TreeNode, sum int) int {
	res=0
	return pathSum3_helper1(root,sum)
}

func pathSum3_helper1(root *Types.TreeNode, sum int) int {
	if root==nil{
		return 0
	}
	pathSum3_helper2(root,sum,&res)
	pathSum3_helper1(root.Left,sum)
	pathSum3_helper1(root.Right,sum)
	return res
}

func pathSum3_helper2(root *Types.TreeNode,sum int,res *int) {
	if root==nil{
		return
	}
	sum-=root.Val
	if sum==0{
		*res+=1
	}
	pathSum3_helper2(root.Left,sum,res)
	pathSum3_helper2(root.Right,sum,res)
}