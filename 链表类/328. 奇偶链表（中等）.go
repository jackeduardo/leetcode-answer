package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{1,2,3,4,5}
	head := Methods.MakeListNode(arr)
	Methods.PrintList(oddEvenList(head))
}
//因为此题每个子结构很难连续操作，所以不用递归
func oddEvenList(head *Types.ListNode) *Types.ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	a:=head
	b:=head.Next
	c:=b
	for a.Next!=nil&&c.Next!=nil{
		a.Next=c.Next//处理奇
		a=a.Next
		c.Next=a.Next//处理偶
		c=c.Next
	}
	a.Next=b //奇尾接偶头
	return head
}