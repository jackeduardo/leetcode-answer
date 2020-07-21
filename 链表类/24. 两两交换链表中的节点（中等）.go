package main

import (
	"leetcode-answer/Methods"
	"leetcode-answer/Types"
)

func main() {
	arr := []int{1,2,3,4,5}
	head := Methods.MakeListNode(arr)
	Methods.PrintList(swapPairs(head))
}
func swapPairs(head *Types.ListNode) *Types.ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	next:=head.Next
	node:=swapPairs(head.Next.Next)
	head.Next.Next=head
	head.Next=node
	return next
}

