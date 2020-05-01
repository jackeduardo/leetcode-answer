package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)
//这就是一个递归的过程，递归的一个非常重要的点就是：不需要去管函数的内部细节是如何处理的，我们只看其函数作用以及输入与输出(这个特点非常重要，理解透了，能够提高非常多的思考效率）
//否则一步一步按照计算机递归过后再退栈的实际操作过程来思考问题非常耗时耗力，而且在理解上很容易犯错误,利用等效转换的思路，思考第一层大的问题，后序的层面依靠递归来解决即可
//所以写递归就是考虑函数作用是什么，递归下去的输入时什么，所以应该先写函数干什么，再写递归位置
func main() {
	arr:="[1,2,5,3,4,null,6]"
	root:=Methods.CreateTree(arr)
	flatten2(root)
	Methods.PrintTree(root)
}
//巧妙利用后序
var pre *Types.TreeNode
func flatten(root *Types.TreeNode)  {
	if root==nil{
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	root.Right=pre
	root.Left=nil
	pre=root

}

//展开法

func flatten2(root *Types.TreeNode){
	if root==nil{
		return
	}
	flatten2(root.Left)
	flatten2(root.Right)
	temp:=root.Right
	root.Right=root.Left
	root.Left=nil
	for root.Right!=nil{
		root=root.Right
	}
	root.Right=temp
}
