package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr:=[]int{1,2,3,4,5}
	head:=Methods.MakeListNode(arr)
	Methods.PrintList(reverseList(head))
}
func reverseList(head *Types.ListNode) *Types.ListNode {
	if head==nil{
		return nil
	}
	if head.Next==nil{
		return head
	}
	node:=reverseList(head.Next)
	head.Next.Next=head
	head.Next=nil
	return node
}
